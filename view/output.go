package view

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/funcs"
	"github.com/xsar/module"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func LiveOutput(name string, watch, interval int64) {
	var index int64
	for {
		line, err := funcs.LivePrint(name)
		if err != nil {
			log.Fatalf("Not Load Module %s", name)
		}
		switch line.(type) {
		case []module.IoMetric:
			LiveMultiOutput(line, index, watch)
		case []module.DfMetric:
			LiveMultiOutput(line, index, watch)
		default:
			LiveSingleOutput(line, index, watch)
		}
		fmt.Println()
		index++
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func Output(name string, watch int64) {
	var metrics module.Metrics
	var bs []byte
	var index int64
	bs, err := ioutil.ReadFile(config.XsarFile)
	if err != nil {
		log.Fatalf("Failed to open %s file", config.XsarFile)
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			os.Exit(0)
		}
		json.Unmarshal(line, &metrics)
		switch name {
		case "cpu":
			SingleOutput(metrics.Cpu, metrics.Now, index, watch)
		case "load":
			SingleOutput(metrics.Load, metrics.Now, index, watch)
		case "mem":
			SingleOutput(metrics.Mem, metrics.Now, index, watch)
		case "tcp":
			SingleOutput(metrics.Tcp, metrics.Now, index, watch)
		case "udp":
			SingleOutput(metrics.Udp, metrics.Now, index, watch)
		case "traffic":
			SingleOutput(metrics.Traffic, metrics.Now, index, watch)
		case "df":
			MultiOutput(metrics.Df, metrics.Now, index, watch)
		case "io":
			MultiOutput(metrics.Io, metrics.Now, index, watch)
		}
		index++
		fmt.Println()
	}
}
