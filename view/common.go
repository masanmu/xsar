package view

import (
	//	"fmt"
	"github.com/xsar/config"
	"strconv"
)

func Format(value interface{}) string {
	switch value.(type) {
	case float64:
		return format(value.(float64))
	case string:
		return value.(string)
	default:
		return "NAN"
	}
}

func format(value float64) string {
	//	fmt.Println(value / config.GiB)
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
