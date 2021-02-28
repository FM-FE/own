package test

import "fmt"

type Interface interface {
	FuncA()
	FuncB()
}

type Default struct {
}

func (d Default) FuncA() {
	fmt.Println("Default func a")
}
func (d Default) FuncB() {
	fmt.Println("Default func b")
}

type A struct {
}

func (a A) FuncA() {
	fmt.Println("A func a")

}
func (a A) FuncB() {
	fmt.Println("A func b")

}

type B struct {
}

func (b B) FuncA() {
	fmt.Println("B func a")
}
func (b B) FuncB() {
	fmt.Println("B func b")
}
