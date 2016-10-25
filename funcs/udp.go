package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"strconv"
	"strings"
)

var udp module.UdpMetric

func UdpMetrics() interface{} {
	return udpMetrics()
}

func udpMetrics() module.UdpMetric {
	content := Open(config.UdpFile)
	value := strings.Split(strings.Split(content, "\n")[9], " ")
	udp.InDatagrams, _ = strconv.ParseFloat(value[1], 64)
	udp.NoPorts, _ = strconv.ParseFloat(value[2], 64)
	udp.InErrors, _ = strconv.ParseFloat(value[3], 64)
	udp.OutDatagrams, _ = strconv.ParseFloat(value[4], 64)

	return udp
}
