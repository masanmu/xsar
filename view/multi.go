package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
)

func MultiHead(value interface{}, now, index int64) error {
	sortLine, err := ConvInterface(value)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}

	SortHead(sortLine)
	fmt.Println()
	index++
	err = FormatTime(now, index%config.MaxList, 1440)
	return nil
}

func MultiValue(value interface{}, now, index int64) error {
	sortLine, err := ConvInterface(value)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}

	values := SortMap(sortLine)
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, "|------------------", config.ColorTag)
	for _, key := range values {
		value := FormatUnit(sortLine[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	fmt.Println()
	return nil
}

func Multi(line interface{}, index, watch, now int64) error {
	err := FormatTime(now, index%config.MaxList, watch)
	if err != nil {
		return errors.New("Over time")
	}
	switch line.(type) {
	case []module.DfMetric:
		metrics := line.([]module.DfMetric)
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				MultiHead(value, now, index)
				break
			}
		}
		fmt.Println()
		for _, value := range metrics {
			MultiValue(value, now, index)
		}

	case []module.IoMetric:
		metrics := line.([]module.IoMetric)
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				MultiHead(value, now, index)
				break
			}
		}
		fmt.Println()
		for _, value := range metrics {
			MultiValue(value, now, index)
		}
	}
	return nil
}
