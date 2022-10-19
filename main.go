package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"MyShield/handlers"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/zserge/lorca"
)

type Response struct {
	http.ResponseWriter
}

func (r *Response) Text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)

	io.WriteString(r, fmt.Sprintf("%s\n", body))
}

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
		handler := http.NewServeMux()
		handler.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
			resp := Response{w}

			resp.Text(http.StatusOK, "hey")
		})
		httpError := http.ListenAndServe(":8080", handler)

		if httpError != nil {
			log.Fatalf(httpError.Error())
		}
		//fmt.Println(discord)
		//fmt.Println(text)
	})

	<-ui.Done()
}
