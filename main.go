package main

import (
	"fmt"
	"log"
	"os"

	"github.com/zserge/lorca"
)

func readFile() string {
	b, err := os.ReadFile("./index.html")
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	return str
}

func main() {
	fmt.Println(readFile())
	ui, err := lorca.New("data:text/html,"+readFile(), "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Code here

	<-ui.Done()
}
