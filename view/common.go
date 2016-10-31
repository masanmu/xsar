package view

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xsar/config"
	"os"
	"sort"
	"strconv"
	"time"
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
		fmt.Printf("%-25s", "Time")
		return nil
	} else {
		tm := time.Unix(unixTime, 0)
		fmt.Printf("%-25s", tm.Format("2006-01-02 15:04:05"))
		return nil
	}
}

func SortHead(dict map[string]interface{}) {
	var head []string
	for key, _ := range dict {
		head = append(head, key)
	}
	sort.Strings(head)

	for _, key := range head {
		fmt.Printf("%12s", key)
	}
}

func SortMap(dict map[string]interface{}) []string {
	var head []string
	for key, _ := range dict {
		head = append(head, key)
	}
	sort.Strings(head)
	return head
}

func ConvInterface(line interface{}) map[string]interface{} {
	content, _ := json.Marshal(line)
	err := json.Unmarshal(content, &line)
	if err != nil {
		os.Exit(-1)
	}

	sortLine, ok := line.(map[string]interface{})
	if !ok {
		os.Exit(-1)
	}

	return sortLine
}
