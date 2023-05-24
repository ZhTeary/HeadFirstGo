package myInterface

import "fmt"

type MyInterface interface {
	MethodWithoutParameters()
	MethodWithParameters(float64 float64)
	MethodWithReturnValue() string
}
type MyType int

func (m MyType) MethodWithoutParameters() {
	fmt.Println("Method Without Parameters Called ")
}

func (m MyType) MethodWithParameters(float float64) {
	fmt.Println("Method With Parameters Called :", float)
}
func (myType MyType) MethodWithReturnValue() string {
	return "Method With Return Value "
}
func (m MyType) MethodNotInInterface() {
	fmt.Println("Method Not In Interface")
}
