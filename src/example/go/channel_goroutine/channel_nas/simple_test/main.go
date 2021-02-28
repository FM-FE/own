package main

import "log"

func main() {
	var ch chan int
	ch = make(chan int)
	go func(ch chan int) {
		ch <- 1
		ch <- 2
	}(ch)

	log.Println(<-ch)

}
