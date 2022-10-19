package main

import (
	"log"

	"MyShield/handlers"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zserge/lorca"
)

func main() {
	var debug = true
	var data = ""
	if debug == true {
		data = handlers.ReadFile("start")
	} else {
		data = handlers.FromGit("start")
	}
	//fmt.Println(data)
	ui, err := lorca.New("data:text/html,"+data, "", 800, 500)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	// Code here
	// Login
	ui.Bind("launch", func() {
		text := ui.Eval("document.getElementById('put').value").String()
		discord, error := discordgo.New("Bot " + text)

		alert := func(msg string) {
			ui.Eval("document.getElementById('alert').innerHTML = '" + msg + "'")
		}

		if error != nil {
			alert("Unknown error occurred.")
			return
		}

		error = discord.Open()
		if error != nil {
			if strings.Contains(error.Error(), "Authentication failed.") {
				alert("Auth failed")
			} else {
				alert("Unknown error occurred.")
			}
			return
		}
		newData := ""
		if debug == true {
			newData = handlers.ReadFile("dashboard")
		} else {
			newData = handlers.FromGit("dashboard")
		}

		ui.Load("data:text/html," + newData)

		handlers.InitWeb()

		handlers.StartUptime(ui)
	})

	<-ui.Done()
}
