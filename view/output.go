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
		sortLine, err := ConvInterface(line)
		if err != nil {
			log.Fatalf("Convert interface to map failed")
		}
		if index == 0 {
			SortHead(sortLine)
		} else {
			err = FormatTime(time.Now().Unix(), index, watch)
			values := SortMap(sortLine)
			for _, key := range values {
				value := FormatUnit(sortLine[key])
				fmt.Printf("%12s", value)
			}

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
		err = FormatTime(metrics.Now, index, watch)
		if err != nil {
			continue
		}
		switch name {
		case "cpu":
			defaultOutput(metrics.Cpu, metrics.Now, index)
		case "load":
			defaultOutput(metrics.Load, metrics.Now, index)
		case "mem":
			defaultOutput(metrics.Mem, metrics.Now, index)
		case "tcp":
			defaultOutput(metrics.Tcp, metrics.Now, index)
		case "udp":
			defaultOutput(metrics.Udp, metrics.Now, index)
		case "traffic":
			defaultOutput(metrics.Traffic, metrics.Now, index)
		case "df":
			multiOutput(metrics.Df, metrics.Now, index, watch)
		case "io":
			multiOutput(metrics.Io, metrics.Now, index, watch)
		}
		index++
		fmt.Println()
	}
}

func defaultOutput(line interface{}, now, index int64) {
	sortLine, err := ConvInterface(line)
	if err != nil {
		log.Fatalf("Convert interface to map failed")
	}

	if index == 0 {
		SortHead(sortLine)
	} else {
		values := SortMap(sortLine)
		for _, key := range values {
			value := FormatUnit(sortLine[key])
			fmt.Printf("%12s", value)
		}
	}
}

func multiOutput(line interface{}, now, index, watch int64) {
	fmt.Println()
	switch line.(type) {
	case []module.DfMetric:
		metrics := line.([]module.DfMetric)
		if index == 0 {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				SortHead(sortLine)
				break
			}
		} else {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				values := SortMap(sortLine)
				for _, key := range values {
					value := FormatUnit(sortLine[key])
					fmt.Printf("%10s", value)
				}
				fmt.Println()
			}
			for i := 0; i < 81; i++ {
				fmt.Printf("%s", "-")
			}
		}

	case []module.IoMetric:
		metrics := line.([]module.IoMetric)
		if index == 0 {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				SortHead(sortLine)
				break
			}
		} else {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				values := SortMap(sortLine)
				for _, key := range values {
					value := FormatUnit(sortLine[key])
					fmt.Printf("%10s", value)
				}
				fmt.Println()
			}
			for i := 0; i < 106; i++ {
				fmt.Printf("%s", "-")
			}
		}
	}
}
