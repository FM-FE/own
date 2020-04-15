package main

import (
	"example/channel_goroutine/channel/ch1"
	"example/channel_goroutine/channel/ch2"
	"log"
)

func main() {
	//{
	//	s := make(chan int)
	//	v := 1
	//	go func() {
	//		<-s
	//		fmt.Println("stop at here")
	//	}()
	//	s <- v
	//	fmt.Println("finish")
	//
	//	var i int
	//	fmt.Println(i)
	//}

	ch := make(chan int)
	go ch1.Ch1(ch)
	log.Println(<-ch)

	go ch2.Ch2(ch)
	log.Println(<-ch)
}
