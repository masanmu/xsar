package main

import (
	"flag"
	"fmt"
	"github.com/xsar/funcs"
	"github.com/xsar/module"
	"github.com/xsar/view"
	"github.com/xsar/xsar"
	"os"
	"strings"
)

func main() {
	var dohelp, cron, listModule bool
	fs := flag.NewFlagSet("xsar", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(Help())
		os.Exit(0)
	}

	fs.BoolVar(&dohelp, "help", false, "")
	fs.BoolVar(&cron, "cron", false, "")
	fs.BoolVar(&listModule, "list", false, "")

	funcs.BuildMappers()
	module.AddCmdFlags(fs)
	fs.Parse(os.Args[1:])

	switch {
	case dohelp:
		fs.Usage()
	case listModule:
		var modules xsar.Module
		modules.ListModule()
	case cron:
		xsar.Colloct()
	default:
		fs.Visit(func(fn *flag.Flag) {
			view.View(fn.Name)
		})
	}
}

func Help() string {
	helpText := `
Usage: tsar [options]
Options:
    -check         display last record for alert
    -check     display last record for alert.example:tsar --check / tsar --check --cpu --io
    -watch     display last records in N mimutes. example:tsar --watch 30 / tsar --watch 30 --cpu --io
    -cron      run in cron mode, output data to file
    -interval  specify intervals numbers, in minutes if with --live, it is in seconds
    -list      list enabled modules
    -live      running print live mode, which module will print
    -file      specify a filepath as input
    -ndays     show the value for the past days (default: 1)
    -date      show the value for the specify day(n or YYYYMMDD)
    -merge     merge multiply item to one
    -detail    do not conver data to K/M/G
    -spec/-s      show spec field data, tsar --cpu -s sys,util
    -item      show spec item data, tsar --io -I sda
    -help      help
`
	helpText += module.Help()
	return strings.TrimSpace(helpText)
}
