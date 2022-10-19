package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FromGit(file string) string {
	resp, err := http.Get("https://raw.githubusercontent.com/dkfrede/MyShield/main/" + file + ".html")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)
}
