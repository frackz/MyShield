package handlers

import (
	"os"
	"log"
)

func ReadFile() string {
	b, err := os.ReadFile("./index.html")
	if err != nil {log.Fatal(err)}
	return string(b)
}