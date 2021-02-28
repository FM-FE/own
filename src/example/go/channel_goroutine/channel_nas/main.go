package main

import (
	"log"
)

func main() {
	var ch chan int
	ch = make(chan int)
	log.Println(">> Start go func Loop")
	nodemap := InitNodemap()
	for _, node := range nodemap {
		go func(n Node, tmp chan int) {
			log.Printf("go func %v", n)
			nodeErr := CheckMount(n, tmp)
			if nodeErr != nil {
				log.Println(nodeErr.Error())
				return
			}
		}(node, ch)
	}

	log.Println("Start receive chan")
	//log.Printf("after receive %v", <-ch)
	//mountStatus := <-ch
	//log.Printf("after evluated %v", mountStatus)
	i := 0
	for tmp := range ch {
		log.Println(tmp)
		i++
		if i == len(nodemap) {
			close(ch)
		}
	}

	log.Println("go routine over")

	return
}
