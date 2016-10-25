package view

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"io/ioutil"
	"log"
	"os"
)

func View(name string) {
	view(name)
}

func view(name string) {
	var metrics module.Metrics
	var bs []byte
	bs, err := ioutil.ReadFile(config.XsarFile)
	if err != nil {
		log.Fatalf("Failed to open %s file", config.XsarFile)
	}
	reader := bufio.NewReader(bytes.NewBuffer(bs))
	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			os.Exit(0)
		}
		json.Unmarshal(line, &metrics)
	}
}
