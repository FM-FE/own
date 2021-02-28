package main

import "example/go/interface/test"

func main() {
	var a test.A
	var b test.B
	var d test.Default
	a.FuncA()
	a.FuncB()
	b.FuncA()
	b.FuncB()
	d.FuncA()
	d.FuncB()
}
