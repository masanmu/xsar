package view

import (
	"fmt"

	"github.com/xsar/config"
)

var maxMetric = map[string]interface{}{}
var minMetric = map[string]interface{}{}
var avgMetric = map[string]interface{}{}
var head []string

func agg(sortHead []string, line map[string]interface{}) {
	head = sortHead
	for _, key := range sortHead {
		maxAgg(key, line[key])
		minAgg(key, line[key])
		avgAgg(key, line[key])
	}
}

func maxAgg(key string, value interface{}) {
	if index == 0 {
		maxMetric[key] = value
	}
	err := ConvInterfaceToFloat(value)
	if err != nil {
		maxMetric[key] = value
		return
	}
	if value.(float64) > (maxMetric[key]).(float64) {
		maxMetric[key] = value
	}
}

func minAgg(key string, value interface{}) {
	if index == 0 {
		minMetric[key] = value
		return
	}
	err := ConvInterfaceToFloat(value)
	if err != nil {
		minMetric[key] = value
		return
	}
	if value.(float64) < minMetric[key].(float64) {
		minMetric[key] = value
	}
}

func avgAgg(key string, value interface{}) {
	if index == 0 {
		avgMetric[key] = value
		return
	}
	err := ConvInterfaceToFloat(value)
	if err != nil {
		avgMetric[key] = value
		return
	}
	avgMetric[key] = avgMetric[key].(float64) + value.(float64)
}

func printAvgAgg() {
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "AVG", config.ColorTag)
	for _, key := range head {
		switch avgMetric[key].(type) {
		case string:
			value := FormatUnit(avgMetric[key])
			fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
		case float64:
			value := FormatUnit(avgMetric[key].(float64) / float64(index))
			fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
		}
	}
	fmt.Println()
}

func printMinAgg() {
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "MIN", config.ColorTag)
	for _, key := range head {
		value := FormatUnit(minMetric[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	fmt.Println()
}
func printMaxAgg() {
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "MAX", config.ColorTag)
	for _, key := range head {
		value := FormatUnit(maxMetric[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	fmt.Println()
}
