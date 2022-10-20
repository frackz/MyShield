package handlers

import (
	"fmt"
	"time"
)

type history struct {
	name string
	by   string
	time int64
}

var his = []history{}

func AddToHistory(name string, by string) {
	bis := append(his, history{name, by, time.Now().Unix()})
	fmt.Printf("bis: %v\n", bis)
}
