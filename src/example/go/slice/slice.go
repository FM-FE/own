package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Tasks []string

func main() {
	Tasks = []string{"01", "02", "03", "04"}
	fmt.Println("print")
	reader := bufio.NewReader(os.Stdin)
	s, e := reader.ReadString('\n')
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Println(s)
	s = strings.Replace(s, "\n", "", -1)
	delnum, e := strconv.Atoi(s)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	for i := 0; i < len(Tasks); i++ {
		if i == delnum {
			Tasks = append(Tasks[0:i], Tasks[i+1:]...)
		}
	}
}
