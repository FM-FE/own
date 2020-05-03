package ch2

import (
	"log"
	"time"
)

func Ch2(a chan int) {
	time.Sleep(10 * time.Second)
	log.Println("2222: channel 2 is running")
	a <- 2
}
