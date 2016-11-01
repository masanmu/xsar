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
	err := FormatTime(time.Now().Unix(), index, watch)
	if err != nil {
		os.Exit(-1)
	}
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
		err = FormatTime(metrics.Now, index%config.MaxList, watch)
		if err != nil {
			continue
		}
		switch name {
		case "cpu":
			SingleOutput(metrics.Cpu, metrics.Now, index)
		case "load":
			SingleOutput(metrics.Load, metrics.Now, index)
		case "mem":
			SingleOutput(metrics.Mem, metrics.Now, index)
		case "tcp":
			SingleOutput(metrics.Tcp, metrics.Now, index)
		case "udp":
			SingleOutput(metrics.Udp, metrics.Now, index)
		case "traffic":
			SingleOutput(metrics.Traffic, metrics.Now, index)
		case "df":
			MultiOutput(metrics.Df, metrics.Now, index, watch)
		case "io":
			MultiOutput(metrics.Io, metrics.Now, index, watch)
		}
		index++
		fmt.Println()
	}
}
