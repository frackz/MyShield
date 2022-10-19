package handlers

import (
	"log"
	"os"
)

func ReadFile(file string) string {
	b, err := os.ReadFile("./html/" + file + ".html")
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}
