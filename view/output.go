package view

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"time"
)

func Output(name string) {
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
			os.Exit(0)
		}
		json.Unmarshal(line, &metrics)
		switch name {
		case "cpu":
			defaultOutput(metrics.Cpu, metrics.Now)
		case "load":
			defaultOutput(metrics.Load, metrics.Now)
		case "mem":
			defaultOutput(metrics.Mem, metrics.Now)
		case "tcp":
			defaultOutput(metrics.Tcp, metrics.Now)
		case "udp":
			defaultOutput(metrics.Udp, metrics.Now)
		case "traffic":
			defaultOutput(metrics.Traffic, metrics.Now)
		case "df":
			for _, df := range metrics.Df {
				defaultOutput(df, metrics.Now)
			}
		case "io":
			for _, io := range metrics.Io {
				defaultOutput(io, metrics.Now)
			}
		}
	}
}

func sortMap(dict map[string]interface{}) []string {
	var head []string
	for key, _ := range dict {
		head = append(head, key)
	}
	sort.Strings(head)
	return head
}

func defaultOutput(line interface{}, now int64) {
	content, _ := json.Marshal(line)
	err := json.Unmarshal(content, &line)
	if err != nil {
		os.Exit(-1)
	}

	tm := time.Unix(now, 0)

	fmt.Printf("%-25s", tm.Format("2006-01-02 15:04:05"))
	sortLine := line.(map[string]interface{})
	head := sortMap(sortLine)
	for _, key := range head {
		value := Format(sortLine[key])
		fmt.Printf("%-10s", value)
	}
	fmt.Println()
}
