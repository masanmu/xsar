package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
)

func SingleOutput(line interface{}, now, watch int64) error {
	sortLine, err := ConvInterfaceToMap(line)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}
	multiValues = nil

	err = FormatTime(now, index%config.MaxList, watch)
	if err != nil {
		index--
		return errors.New("Over time")
	}

	if index%config.MaxList == 0 {
		PrintHead(sortLine)
		FormatTime(now, (index+1)%config.MaxList, watch)
	}
	sortHead = SortMap(sortLine)
	multiAgg(sortHead, append(multiValues, sortLine))
	for _, key := range head {
		value := FormatUnit(sortLine[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	fmt.Println()
	return nil
}

func MultiOutput(line interface{}, now, watch int64) error {
	err := Multi(line, watch, now)
	return err
}
