package main

import "fmt"

func main() {
	s := make(chan int)
	v := 1
	go func() {
		<-s
		fmt.Println("stop at here")
	}()
	s <- v
	fmt.Println("finish")
	var i int
	fmt.Println(i)
}
