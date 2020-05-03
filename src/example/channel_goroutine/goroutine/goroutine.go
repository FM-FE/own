package main

import (
	"fmt"
	"time"
)

func main() {

	var a int
	a = 1
	fmt.Print("aaaa")
	go func() {
		time.Sleep(5 * time.Millisecond)
		a++
	}()

	go func() {
		time.Sleep(2 * time.Millisecond)
		a++
	}()

	go func() {
		time.Sleep(1 * time.Millisecond)
		a++
	}()

	for {
		if a == 1 {
			fmt.Println("ONE")
		}

		if a == 2 {
			fmt.Println("TWO")
		}

		if a == 3 {
			fmt.Println("THREE")
		}
		if a == 4 {
			fmt.Print("FOUR")
		}
	}
}
