package main

import (
	"example/go/init/subinit"
	"log"
)

func main() {
	log.Println("In main")
}

func init() {
	log.Println("In init")
	subinit.CallSubInit()
}
