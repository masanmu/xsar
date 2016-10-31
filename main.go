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
	var watch int64
	fs := flag.NewFlagSet("xsar", flag.ExitOnError)
	fs.Usage = func() {
		fmt.Println(Help())
		os.Exit(0)
	}

	fs.BoolVar(&dohelp, "help", false, "")
	fs.BoolVar(&cron, "c", false, "")
	fs.Int64Var(&watch, "w", 1440, "")

	funcs.BuildMappers()
	module.AddCmdFlags(fs)
	fs.Parse(os.Args[1:])

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
    -c      run in cron mode, output data to file
    -w     display last records in N mimutes.
`
	helpText += module.Help()
	return strings.TrimSpace(helpText)
}
