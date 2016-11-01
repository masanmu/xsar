package view

import (
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"log"
)

func SingleOutput(line interface{}, now, index int64) {
	sortLine, err := ConvInterface(line)
	if err != nil {
		log.Fatalf("Convert interface to map failed")
	}

	if index%config.MaxList == 0 {
		SortHead(sortLine)
	} else {
		values := SortMap(sortLine)
		for _, key := range values {
			value := FormatUnit(sortLine[key])
			fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
		}
	}
}

func MultiOutput(line interface{}, now, index, watch int64) {
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
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				values := SortMap(sortLine)
				for _, key := range values {
					value := FormatUnit(sortLine[key])
					fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
				}
				fmt.Println()
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
			for _, value := range metrics {
				sortLine, err := ConvInterface(value)
				if err != nil {
					log.Fatalf("Convert interface to map failed")
				}

				values := SortMap(sortLine)
				for _, key := range values {
					value := FormatUnit(sortLine[key])
					fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
				}
				fmt.Println()
			}
		}
	}
}
