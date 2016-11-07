package view

import ()

var index int64
var sortHead []string
var multiValues = []map[string]interface{}{}

func View(name string, watch int64) {
	Output(name, watch)
}

func Live(name string, watch, interval int64) {
	LiveOutput(name, watch, interval)
}
