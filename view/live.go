package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
	"time"
)

func LiveSingleOutput(line interface{}, watch int64) error {
	sortLine, err := ConvInterface(line)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}
	if index%config.MaxList == 0 {
		err = FormatTime(time.Now().Unix(), index%config.MaxList, watch)
		SortHead(sortLine)
		fmt.Println()
	}
	err = FormatTime(time.Now().Unix(), index+1, watch)
	if err != nil {
		return errors.New("Over time")
	}
	values := SortMap(sortLine)
	for _, key := range values {
		value := FormatUnit(sortLine[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	return nil
}

func LiveMultiOutput(line interface{}, watch int64) error {
	var now int64 = time.Now().Unix()
	err := Multi(line, watch, now)
	return err
}
