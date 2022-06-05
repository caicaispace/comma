package test

import "fmt"

type variable struct{}

var v *variable

func getVariableInstance() *variable {
	if v == nil {
		fmt.Println("-------------- getVariableInstance --------------")
		v = new(variable)
	}
	return v
}

func (v *variable) variableFun() {
	fmt.Println(v)
}
