package xsar

import (
	"encoding/json"
	"github.com/xsar/config"
	"github.com/xsar/funcs"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var writeLine string

func Colloct() {
	now := time.Now()
	writeLine = "{\"" + "now\":" + strconv.FormatInt(now.Unix(), 10) + ","
	for _, v := range funcs.Mappers {
		colloct(v.Name, v.Fs())
	}
	writeFile(strings.TrimRight(writeLine, ",") + "}\n")
}

func colloct(name string, item interface{}) {
	content, err := json.Marshal(item)
	if err != nil {
		log.Fatalf("%s", err)
	}
	info := "\"" + name + "\":" + string(content) + ","
	writeLine += info
}

func writeFile(content string) {
	f, err := os.OpenFile(config.XsarFile, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer f.Close()
	f.WriteString(content)
}
