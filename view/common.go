package view

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/xsar/config"
)

func FormatUnit(value interface{}) string {
	switch value.(type) {
	case float64:
		return formatUnit(value.(float64))
	case string:
		return value.(string)
	default:
		return "NAN"
	}
}

func formatUnit(value float64) string {
	if value/config.TiB > 1 {
		return strconv.FormatFloat(value/config.TiB, 'f', 2, 64) + "T"
	} else if value/config.GiB > 1 {
		return strconv.FormatFloat(value/config.GiB, 'f', 2, 64) + "G"
	} else if value/config.MiB > 1 {
		return strconv.FormatFloat(value/config.MiB, 'f', 2, 64) + "G"
	} else if value/config.KiB > 1 {
		return strconv.FormatFloat(value/config.KiB, 'f', 2, 64) + "K"
	} else {
		return strconv.FormatFloat(value, 'f', 2, 64)
	}
}

func FormatTime(unixTime, index, watch int64) error {
	now := time.Now().Unix()
	if now-unixTime >= watch*60 {
		return errors.New("Over time")
	}
	if index == 0 {
		fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "Time", config.ColorTag)
		return nil
	} else {
		tm := time.Unix(unixTime, 0)
		fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, tm.Format("2006-01-02 15:04:05"), config.ColorTag)
		return nil
	}
}

func PrintHead(dict map[string]interface{}) {
	head := SortMap(dict)
	for _, key := range head {
		fmt.Printf(config.FormatHeadString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, key, config.ColorTag)
	}
	fmt.Println()
}

func SortMap(dict map[string]interface{}) []string {
	var head []string
	for key, _ := range dict {
		head = append(head, key)
	}
	sort.Strings(head)
	return head
}

func ConvInterfaceToMap(line interface{}) (map[string]interface{}, error) {
	content, _ := json.Marshal(line)
	err := json.Unmarshal(content, &line)
	if err != nil {
		return nil, errors.New("Unmarshal json failed")
	}

	sortLine, ok := line.(map[string]interface{})
	if !ok {
		return nil, errors.New("Conv Interface to map error")
	}

	return sortLine, nil
}

func ConvInterfaceToFloat(value interface{}) error {
	switch value.(type) {
	case float64:
		return nil
	case string:
		return errors.New("string")
	default:
		return errors.New("default")
	}
}

func initVariables(len int) {
	maxMetric = make([]map[string]interface{}, len)
	minMetric = make([]map[string]interface{}, len)
	avgMetric = make([]map[string]interface{}, len)

	for i, _ := range maxMetric {
		maxMetric[i] = map[string]interface{}{}
	}
	for i, _ := range minMetric {
		minMetric[i] = map[string]interface{}{}
	}
	for i, _ := range avgMetric {
		avgMetric[i] = map[string]interface{}{}
	}
}
