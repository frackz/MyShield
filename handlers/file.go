package handlers

import (
	"os"
	"log"
)

func ReadFile() string {
	b, err := os.ReadFile("./html/index.html")
	if err != nil {log.Fatal(err)}
	return string(b)
}