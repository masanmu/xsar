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
	for {
		line, err := funcs.LivePrint(name)
		if err != nil {
			log.Fatalf("Not Load Module %s", name)
		}
		switch line.(type) {
		case []module.IoMetric:
			LiveMultiOutput(line, watch)
		case []module.DfMetric:
			LiveMultiOutput(line, watch)
		default:
			LiveSingleOutput(line, watch)
		}
		fmt.Println()
		index++
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func Output(name string, watch int64) {
	var metrics module.Metrics
	var bs []byte
	bs, err := ioutil.ReadFile(config.XsarFile)
	if err != nil {
		log.Fatalf("Failed to open %s file", config.XsarFile)
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			fmt.Println()
			printAvgAgg()
			printMaxAgg()
			printMinAgg()
			os.Exit(0)
		}
		json.Unmarshal(line, &metrics)
		switch name {
		case "cpu":
			SingleOutput(metrics.Cpu, metrics.Now, watch)
		case "load":
			SingleOutput(metrics.Load, metrics.Now, watch)
		case "mem":
			SingleOutput(metrics.Mem, metrics.Now, watch)
		case "tcp":
			SingleOutput(metrics.Tcp, metrics.Now, watch)
		case "udp":
			SingleOutput(metrics.Udp, metrics.Now, watch)
		case "traffic":
			SingleOutput(metrics.Traffic, metrics.Now, watch)
		case "df":
			MultiOutput(metrics.Df, metrics.Now, watch)
		case "io":
			MultiOutput(metrics.Io, metrics.Now, watch)
		}
		index++
	}

}
