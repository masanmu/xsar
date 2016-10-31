package view

import ()

func View(name string, watch int64) {
	Output(name, watch)
}

func Live(name string, watch, interval int64) {
	LiveOutput(name, watch, interval)
}
