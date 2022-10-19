package main

import (
	//"fmt"
	"log"

	"MyShield/handlers"
	"github.com/zserge/lorca"
)

func main() {

	var debug = true

	var data = ""

	if debug == true {
		data = handlers.ReadFile()
	} else {
		data = handlers.FromGit()
	}

	ui, err := lorca.New("data:text/html,"+data, "", 480, 320)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Code here

	<-ui.Done()
}
