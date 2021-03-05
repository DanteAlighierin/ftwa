package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("test" "e")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test" "e", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
