package ch2

import "log"

func Ch2(a chan int) {
	log.Println("2222: channel 2 is running")
	a <- 2
}
