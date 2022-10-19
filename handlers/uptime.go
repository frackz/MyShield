package handlers

import (
	"strconv"
	"time"

	"github.com/zserge/lorca"
)

var gui lorca.UI

var days int
var hours int
var minutes int
var seconds int

func StartUptime(ui lorca.UI) {
	gui = ui
	for true {
		if seconds == 59 {
			minutes += 1
			seconds = 0
			if minutes == 60 {
				minutes = 0
				hours += 1
				if hours == 24 {
					hours = 0
					days += 1
				}
			}
		} else {
			seconds += 1
		}
		ui.Eval("document.getElementById('uptime').innerHTML = 'Uptime: " + strconv.Itoa(days) + " day(s) " + strconv.Itoa(hours) + " hour(s) " + strconv.Itoa(minutes) + " minute(s) " + strconv.Itoa(seconds) + " second(s)'")
		time.Sleep(time.Second * 1)
	}
}
