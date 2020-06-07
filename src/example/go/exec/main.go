package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	command := exec.Command("sh", "-c", "lk")
	output, err := command.CombinedOutput()
	fmt.Println(err.Error())
	fmt.Println(string(output))

	cmdStr := fmt.Sprintf("cat /etc/mtab| grep /var/lib/docker/containers/d82a322f684b97725f9c97d01c6de8d5a")
	cmd := exec.Command("sh", "-c", cmdStr)
	output, e := cmd.CombinedOutput()
	if e != nil {
		if strings.Contains(e.Error(), "exit status 1") && string(output) == "" {
			fmt.Println("OK")
			return
		}
		fmt.Println(e.Error())
		return
	}
	fmt.Println(string(output))
}
