package main

import (
	"example/go/channel_goroutineine/channel/ch1"
	"example/go/channel_goroutineine/channel/ch2"
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

	c := make(chan int)
	//var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		if i == 0 {
			//	wg.Add(1)
			go ch1.Ch1(c)
		}
		if i == 1 {
			//wg.Add(1)
			go ch2.Ch2(c)
		}
	}

	for i := 0; i < 2; i++ {
		log.Println("after go routine")
		log.Println(<-c)
	}

	//time.Sleep(time.Millisecond)
	//close(c)

	//go func() {
	//	time.Sleep(1 * time.Hour)
	//}()
	//c := make(chan int)
	//go func() {
	//	for i := 0; i < 10; i = i + 1 {
	//
	//		c <- i
	//	}
	//	close(c) //如果将此句注释掉，那么下面的for range在打印完管道的内容后会一直阻塞。
	//}()
	//for i := range c {
	//	fmt.Println(i)
	//}
	//fmt.Println("Finished")

}
