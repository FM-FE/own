package main

import (
	"example/go/os/env"
	"log"
	"os"
	"time"
)

func main() {
	//file, e := os.OpenFile("/root/test", os.O_RDWR|os.O_CREATE, 0644)
	//if e != nil {
	//	return
	//}
	//log.Println(file)
	//_, e = os.Stat("/root/test")
	//if e == nil {
	//	log.Println("file exist")
	//}

	//_, e := os.Stat("/root/dasd")
	//
	//if e != nil && strings.Contains(e.Error(), "no such file or directory") {
	//	log.Println("catch error")
	//}

	go func() {
		e := os.Setenv("GDASURL", "https://test")
		if e != nil {
			log.Println(e.Error())
			return
		}
		log.Println("go func end")
	}()

	time.Sleep(2 * time.Second)
	gdasbygetenv := os.Getenv("GDASURL")
	log.Println("gdasbygetenv: " + gdasbygetenv)
	log.Println("gdasfrompackage" + env.GdasURL)

	url := env.GetURL()
	log.Println("env.GetURL()", url)
}
