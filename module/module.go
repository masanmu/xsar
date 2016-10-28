package module

import (
	"flag"
)

type ModuleConfig struct {
	Cpu       bool
	Mem       bool
	Swap      bool
	Tcp       bool
	Udp       bool
	Traffic   bool
	Io        bool
	Pcsw      bool
	Partition bool
	Tcpx      bool
	Load      bool
}

var module ModuleConfig

func AddCmdFlags(fs *flag.FlagSet) {
	fs.BoolVar(&module.Cpu, "cpu", false, "CPU share (user, system, interrupt, nice, & idle)")
	fs.BoolVar(&module.Mem, "mem", false, "Physical memory share (active, inactive, cached, free, wired)")
	fs.BoolVar(&module.Swap, "swap", false, "swap usage")
	fs.BoolVar(&module.Tcp, "tcp", false, "TCP traffic     (v4)")
	fs.BoolVar(&module.Udp, "udp", false, "udP traffic     (v4)")
	fs.BoolVar(&module.Traffic, "traffic", false, "Net traffic statistics")
	fs.BoolVar(&module.Io, "io", false, "Linux I/O performance")
	fs.BoolVar(&module.Pcsw, "pcsw", false, "Process (task) creation and context switch")
	fs.BoolVar(&module.Partition, "partition", false, "Disk and partition usage")
	fs.BoolVar(&module.Tcpx, "tcpx", false, "TCP connection data")
	fs.BoolVar(&module.Load, "load", false, "System Run Queue and load average")
}

func Help() string {
	helpText := `
Modules Enabled:
    --cpu               CPU share (user, system, interrupt, nice, & idle)
    --mem               Physical memory share (active, inactive, cached, free, wired)
    --tcp               TCP traffic     (v4)
    --udp               UDP traffic     (v4)
    --traffic           Net traffic statistics
    --io                Linux I/O performance
    --df         Disk and partition usage
    --load              System Run Queue and load average
`
	return helpText
}
