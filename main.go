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
Usage: xsar [options]
Options:
    -cron      run in cron mode, output data to file
    -interval  specify intervals numbers, in minutes if with --live, it is in seconds
    -list      list enabled modules
`
	helpText += module.Help()
	return strings.TrimSpace(helpText)
}
