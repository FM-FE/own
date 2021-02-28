package main

import "strconv"

type Node struct {
	IP   string
	Name string
}

func InitNodemap() (nodemap []Node) {
	node1 := Node{
		IP:   "1",
		Name: "one",
	}
	node2 := Node{
		IP:   "2",
		Name: "two",
	}
	node3 := Node{
		IP:   "3",
		Name: "three",
	}
	nodemap = append(nodemap, node1, node2, node3)
	return
}

func CheckMount(node Node, ch chan int) (e error) {
	i, e := strconv.Atoi(node.IP)
	if e != nil {
		return e
	}
	ch <- i
	return
}
