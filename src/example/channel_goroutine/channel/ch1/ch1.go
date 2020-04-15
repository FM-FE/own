package ch1

import "log"

func Ch1(a chan int) {
	a <- 1
	log.Println("1111: channel 1 is running")

}
