package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
)

var maxMetric = []map[string]interface{}{}
var minMetric = []map[string]interface{}{}
var avgMetric = []map[string]interface{}{}
var head []string

func agg(sortHead []string, line map[string]interface{}, cur int) {
	for _, key := range sortHead {
		err := preAgg(key, line[key], cur)
		if err != nil {
			continue
		}
		maxAgg(key, line[key], cur)
		minAgg(key, line[key], cur)
		avgAgg(key, line[key], cur)
	}
}

func multiAgg(sortHead []string, values []map[string]interface{}) {
	head = sortHead
	if index == 0 {
		len := len(values)
		initVariables(len)
	}
	for cur, key := range values {
		agg(sortHead, key, cur)
	}
}

func maxAgg(key string, value interface{}, cur int) {
	if value.(float64) > maxMetric[cur][key].(float64) {
		maxMetric[cur][key] = value
	}
}

func minAgg(key string, value interface{}, cur int) {
	if value.(float64) < minMetric[cur][key].(float64) {
		minMetric[cur][key] = value
	}
}

func avgAgg(key string, value interface{}, cur int) {
	avgMetric[cur][key] = avgMetric[cur][key].(float64) + value.(float64)
}

func preAgg(key string, value interface{}, cur int) error {
	if index == 0 {
		avgMetric[cur][key] = value
		minMetric[cur][key] = value
		maxMetric[cur][key] = value
		return errors.New("init map")
	}
	err := ConvInterfaceToFloat(value)
	if err != nil {
		avgMetric[cur][key] = value
		minMetric[cur][key] = value
		maxMetric[cur][key] = value
		return errors.New("not float")
	}
	return nil
}

func printAgg(flag string, metrics []map[string]interface{}) {
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, flag, config.ColorTag)
	for _, values := range metrics {
		fmt.Println()
		fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, config.FormatValueHeadString, config.ColorTag)
		for _, key := range head {
			if flag == "AVG" {
				switch values[key].(type) {
				case string:
					value := FormatUnit(values[key])
					fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
				case float64:
					value := FormatUnit(values[key].(float64) / float64(index))
					fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
				}
			} else {
				value := FormatUnit(values[key])
				fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
			}
		}
	}
	fmt.Println()
}
