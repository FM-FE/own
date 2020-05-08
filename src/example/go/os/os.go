package main

import (
	"log"
	"os"
)

func main() {
	file, e := os.OpenFile("/root/test", os.O_RDWR|os.O_CREATE, 0644)
	if e != nil {
		return
	}
	log.Println(file)
	_, e = os.Stat("/root/test")
	if e == nil {
		log.Println("file exist")
	}
}
