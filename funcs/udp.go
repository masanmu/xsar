package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"strconv"
	"strings"
	"time"
)

var udpNow, udpPre, udp module.UdpMetric

func UdpMetrics() interface{} {
	udpPre = udpMetrics()
	time.Sleep(time.Second)
	udpNow = udpMetrics()
	return udpAvg()
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

func udpAvg() module.UdpMetric {
	udp.InDatagrams = Delta(udpNow.InDatagrams, udpPre.InDatagrams)
	udp.InErrors = Delta(udpNow.InErrors, udpPre.InErrors)
	udp.NoPorts = Delta(udpNow.NoPorts, udpPre.NoPorts)
	udp.OutDatagrams = Delta(udpNow.OutDatagrams, udpPre.OutDatagrams)
	return udp
}
