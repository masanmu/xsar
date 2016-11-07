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
	var dohelp, cron bool
	var live string
	var watch, interval int64

	fs := flag.NewFlagSet("xsar", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(Help())
		os.Exit(0)
	}

	fs.BoolVar(&dohelp, "help", false, "")
	fs.BoolVar(&cron, "c", false, "")
	fs.StringVar(&live, "l", "", "")
	fs.Int64Var(&watch, "w", 1440, "")
	fs.Int64Var(&interval, "i", 1, "")

	funcs.BuildMappers()
	module.AddCmdFlags(fs)

	if (len(os.Args[1:])) <= 0 {
		fs.Usage()
	}

	fs.Parse(os.Args[1:])

	if len(live) > 0 {
		view.Live(live, watch, interval)
		os.Exit(0)
	}
	switch {
	case dohelp:
		fs.Usage()
	case cron:
		xsar.Colloct()
	default:
		fs.Visit(func(fn *flag.Flag) {
			view.View(fn.Name, watch)
		})
	}
}

func Help() string {
	helpText := `
Usage: xsar [options]
Options:
    -c/cron      run in cron mode, output data to file
    -w/watch     display last records in N mimutes.
    -l/live	 running print live mode, which module will print
    -i/interval  specify intervals numbers, in minutes if with --live, it is in seconds
`
	helpText += module.Help()
	return strings.TrimSpace(helpText)
}
