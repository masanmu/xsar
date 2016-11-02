package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
)

func SingleOutput(line interface{}, now, index, watch int64) error {
	sortLine, err := ConvInterface(line)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}

	err = FormatTime(now, index%config.MaxList, watch)
	if err != nil {
		return errors.New("Over time")
	}

	if index%config.MaxList == 0 {
		SortHead(sortLine)
		fmt.Println()
		index++
		err = FormatTime(now, index%config.MaxList, watch)
	}
	values := SortMap(sortLine)
	for _, key := range values {
		value := FormatUnit(sortLine[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	return nil
}

func MultiOutput(line interface{}, now, index, watch int64) error {
	err := Multi(line, index, watch, now)
	return err
}
