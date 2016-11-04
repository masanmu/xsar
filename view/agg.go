package view

import (
	"fmt"
	"github.com/xsar/config"
)

var maxMetric = map[string]float64{}
var minMetric = map[string]float64{}
var avgMetric = map[string]float64{}
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
	maxValue, err := ConvInterfaceToFloat(value)
	if err != nil {
		maxMetric[key] = maxValue
		return
	}
	if maxValue > maxMetric[key] {
		maxMetric[key] = maxValue
	}
}

func minAgg(key string, value interface{}) {
	minValue, err := ConvInterfaceToFloat(value)
	if index == 0 {
		minMetric[key] = minValue
		return
	}
	if err != nil {
		minMetric[key] = minValue
		return
	}
	if minValue < minMetric[key] {
		minMetric[key] = minValue
	}
}

func avgAgg(key string, value interface{}) {
	avgValue, err := ConvInterfaceToFloat(value)
	if err != nil {
		avgMetric[key] = avgValue
		return
	}
	avgMetric[key] += avgValue
}

func printAvgAgg() {
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "AVG", config.ColorTag)
	for _, key := range head {
		value := FormatUnit(avgMetric[key] / float64(index))
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
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
