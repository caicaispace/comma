package util

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/goinggo/mapstructure"
)

// a, b := 2, 3
// max := If(a > b, a, b).(int)
func If(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}

func StructToMap(in interface{}, tagName string) (map[string]interface{}, error) {
	t := reflect.TypeOf(in)
	v := reflect.ValueOf(in)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("ToMap only accepts struct or struct pointer; got %T", v)
	}
	out := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		tagName := getStructTag(t.Field(i), tagName)
		if tagName == "" {
			continue
		}
		out[tagName] = v.Field(i).Interface()
	}
	return out, nil
}

func Struct2Json(jsonStr string, obj *interface{}) {
	json.Unmarshal([]byte(jsonStr), &obj)
}

func MapToStruct(inData map[string]interface{}, obj interface{}) error {
	mapData := map[string]interface{}{}
	for k, v := range inData {
		mapData[Case2Camel(k)] = v
	}
	if err := mapstructure.Decode(mapData, obj); err != nil {
		return err
	}
	return nil
}

func getStructTag(f reflect.StructField, tagName string) string {
	return string(f.Tag.Get(tagName))
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.Title(name)
	return strings.ReplaceAll(name, " ", "")
}
