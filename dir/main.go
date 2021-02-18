package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Stat("test")
	if os.IsNotExist(err) {
		errDir := os.MkdirAll("test", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
	}
}
