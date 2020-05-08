package subinit

import (
	"log"
)

func init() {
	log.Println("in subinit")
}

func CallSubInit() {
	log.Println("call subinit")
}
