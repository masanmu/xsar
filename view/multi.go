package view

import (
	"errors"
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
)

func MultiHead(value interface{}, now int64) error {
	sortLine, err := ConvInterfaceToMap(value)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}

	PrintHead(sortLine)
	err = FormatTime(now, (index+1)%config.MaxList, 1440)
	return nil
}

func MultiValue(value interface{}, now int64) error {
	sortLine, err := ConvInterfaceToMap(value)
	if err != nil {
		return errors.New("Convert interface to map failed")
	}
	multiValues = append(multiValues, sortLine)
	sortHead = SortMap(sortLine)
	fmt.Printf(config.FormatTimeString, config.ColorTag, config.Flag, config.BackGround, config.Prospect, config.FormatValueHeadString, config.ColorTag)
	for _, key := range sortHead {
		value := FormatUnit(sortLine[key])
		fmt.Printf(config.FormatValueString, config.ColorTag, value, config.ColorTag)
	}
	fmt.Println()
	return nil
}

func Multi(line interface{}, watch, now int64) error {
	err := FormatTime(now, index%config.MaxList, watch)
	if err != nil {
		index--
		return errors.New("Over time")
	}
	multiValues = nil
	switch line.(type) {
	case []module.DfMetric:
		metrics := line.([]module.DfMetric)
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				MultiHead(value, now)
				break
			}
		}
		fmt.Println()
		for _, value := range metrics {
			MultiValue(value, now)
		}

	case []module.IoMetric:
		metrics := line.([]module.IoMetric)
		if index%config.MaxList == 0 {
			for _, value := range metrics {
				MultiHead(value, now)
				break
			}
		}
		fmt.Println()
		for _, value := range metrics {
			MultiValue(value, now)
		}
	}
	multiAgg(sortHead, multiValues)
	return nil
}
