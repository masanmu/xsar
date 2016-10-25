package funcs

import ()

type MetricValue struct {
	Metric string
	Value  string
}

type Funcs struct {
	Fs   func() interface{}
	Name string
}

var Mappers []Funcs

func BuildMappers() {
	Mappers = []Funcs{
		Funcs{
			Fs:   LoadMetrics,
			Name: "load",
		},
		Funcs{
			Fs:   CpuMetrics,
			Name: "cpu",
		},
		Funcs{
			Fs:   IoMetrics,
			Name: "io",
		},
		Funcs{
			Fs:   TcpMetrics,
			Name: "tcp",
		},
		Funcs{
			Fs:   TrafficMetrics,
			Name: "traffic",
		},
		Funcs{
			Fs:   UdpMetrics,
			Name: "udp",
		},
		Funcs{
			Fs:   MemMetrics,
			Name: "mem",
		},
		Funcs{
			Fs:   DfMetrics,
			Name: "df",
		},
	}
}
