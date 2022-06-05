package test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	title := "4k地球旋转,,,<>;/;《，。，。！@#￥%……&*（）——+{}【】~·!@#$%^&*()_+[]{}<>/*-+很快"
	re := regexp.MustCompile(`[\p{P}\p{S}\n]`) // 去除换行符以及所有符号
	// re := regexp.MustCompile(`([#@])|[\p{P}\p{S}\d]`) // 替换除＃和@以外的所有符号
	// re := regexp.MustCompile(`[\p{P}\p{S}\s\n\d]`) // 去除空格、去除换行符以及所有符号
	title = re.ReplaceAllString(title, "$1")
	fmt.Println(title)
}

func TestStringReplace(t *testing.T) {
	str := "/gateway/search_all_v2/all/_search"
	str = strings.Replace(str, "/gateway", "", 1)
	fmt.Println(str)
}

func TestPointVar(t *testing.T) {
	var v string
	for i := 0; i < 10; i++ {
		pointVar(&v)
		// fmt.Println(v)
		t.Log(v)
	}
}

func pointVar(variable *string) {
	// fmt.Println(len(vars[0].(string)))
	vvv := "12212"
	variable = &vvv
	// vars[0] = util.GetRandomString(10)
}

func TestVar(t *testing.T) {
	fmt.Println(getVariableInstance())
	fmt.Println(getVariableInstance())
}

func TestForArrValuePoint(t *testing.T) {
	arr := []string{
		"1",
		"2",
		"3",
	}
	for i := range arr {
		arr[i] = "888"
	}
	fmt.Println(arr)
}

func TestArrDefaultValue(t *testing.T) {
	var synonymsSlice []*map[string][]string
	maxProId := 2
	for i := 0; i < maxProId+1; i++ {
		synonymsSlice = append(synonymsSlice, nil)
	}
	fmt.Println(synonymsSlice)

	synonymsSlice2 := make([]*map[string][]string, maxProId+1)
	fmt.Println(synonymsSlice2)
}

func TestVariable(t *testing.T) {
	var baseDistance float32
	t.Log(baseDistance)
}
