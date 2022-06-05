package util

import (
	"fmt"
	"goaway/pkg/library/util"
	"testing"
)

func TestIf(t *testing.T) {
	tests := []struct {
		env  string
		want string
	}{
		// TODO: Add test cases.
		{
			env:  "dev",
			want: "127.0.0.1",
		},
		{
			env:  "pro",
			want: "localhost",
		},
	}
	for _, tt := range tests {
		t.Run(tt.env, func(t *testing.T) {
			host := util.If(tt.env == "dev", "127.0.0.1", "localhost").(string)
			if tt.env == "dev" {
				if host != tt.want {
					t.Log(host)
					t.Errorf("got %s, want %s", host, tt.want)
				}
			}
			if tt.env == "pro" {
				if host != tt.want {
					t.Log(host)
					t.Errorf("got %s, want %s", host, tt.want)
				}
			}
		})
	}
}

type structToMap struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestStructToMap(t *testing.T) {
	structToMap := structToMap{
		Id:   1,
		Name: "test",
	}
	result, _ := util.StructToMap(structToMap, "json")
	fmt.Println(result)
}

type PersonStruct struct {
	Name string
	Age  uint8
}

func TestMapToStruct(t *testing.T) {
	mapData := map[string]interface{}{
		"name": "test",
		"age":  18,
	}
	person := PersonStruct{}
	err := util.MapToStruct(mapData, &person)
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Log(person)
}

func TestCase2Camel(t *testing.T) {
	mapData := map[string]interface{}{
		"name_key": "test",
		"age_key":  18,
	}
	for k := range mapData {
		t.Log(util.Case2Camel(k))
	}
}
