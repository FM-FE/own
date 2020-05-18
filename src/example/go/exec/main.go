package main

import (
	"fmt"
	"os/exec"
)

func main() {
	command := exec.Command("sh", "-c", "lk")
	output, err := command.CombinedOutput()
	fmt.Println(err.Error())
	fmt.Println(string(output))
}
