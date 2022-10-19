package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Response struct {
	http.ResponseWriter
}

func (r *Response) Text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)
	io.WriteString(r, fmt.Sprintf("%s\n", body))
}

func InitWeb() {
	http.HandleFunc("/get/", func(w http.ResponseWriter, r *http.Request) {
		resp := Response{w}
		resp.Text(http.StatusOK, "hey")
	})
	httpError := http.ListenAndServe(":8080", nil)
	if httpError != nil {
		log.Fatalf(httpError.Error())
	}
	return
}
