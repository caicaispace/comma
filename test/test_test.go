package test

import (
	"fmt"
	"regexp"
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
