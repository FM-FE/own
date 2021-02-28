package env

import (
	"log"
	"os"
)

var GdasURL = getenv()

func getenv() (s string) {
	return os.Getenv("GDASURL")

}

func GetURL() (url string) {
	log.Println("os.Getenv : " + os.Getenv("GDASURL"))
	log.Println("GdasURL is : " + GdasURL)
	url = GdasURL
	return
}
