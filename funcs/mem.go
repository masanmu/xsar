package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"strconv"
	"strings"
)

var mem module.MemMetric

func MemMetrics() interface{} {
	return memMetrics()
}

func memMetrics() module.MemMetric {
	var kv = map[string]float64{}

	content := strings.TrimRight(Open(config.MemFile), "\n")
	values := strings.Split(content, "\n")

	for _, i := range values {
		metric := strings.TrimRight(strings.Replace(i, " ", "", -1), "kB")
		key := strings.Split(metric, ":")[0]
		value, _ := strconv.ParseFloat(strings.Split(metric, ":")[1], 64)
		kv[key] = value
	}

	mem.MemTotal = kv["MemTotal"]
	mem.MemFree = kv["MemFree"]
	mem.MemUsed = kv["MemTotal"] - kv["MemFree"] - kv["Buffers"] - kv["Cached"]
	mem.Buffers = kv["Buffers"]
	mem.Cached = kv["Cached"]
	mem.SwapTotal = kv["SwapTotal"]
	mem.SwapFree = kv["SwapFree"]
	mem.SwapUsed = kv["MemTotal"] - kv["MemFree"] - kv["Buffers"] - kv["Cached"]

	return mem
}
