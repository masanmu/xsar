package view

import (
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"log"
	"time"
)

func LiveSingleOutput(line interface{}, index, watch int64) {
	sortLine, err := ConvInterface(line)
	if err != nil {
		log.Fatalf("Convert interface to map failed")
	}
	if index%config.MaxList == 0 {
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
}

func LiveMultiOutput(line interface{}, index, watch int64) {
	fmt.Println()
	switch line.(type) {
	case []module.DfMetric:
		metrics := line.([]module.DfMetric)
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				SortHead(sortLine)
				break
			}
		} else {
			err := FormatTime(time.Now().Unix(), index, watch)
			if err != nil {
				log.Fatalf("Over time")
			}
			fmt.Println()
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
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				SortHead(sortLine)
				break
			}
		} else {
			err := FormatTime(time.Now().Unix(), index, watch)
			if err != nil {
				log.Fatalf("Over time")
			}
			fmt.Println()
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
