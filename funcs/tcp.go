package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"strconv"
	"strings"
	"time"
)

var tcpNow, tcpPre, tcp module.TcpMetric

func TcpMetrics() interface{} {
	tcpPre = tcpMetrics()
	time.Sleep(time.Second)
	tcpNow = tcpMetrics()
	return tcpAvg()
}

func tcpMetrics() module.TcpMetric {
	content := Open(config.TcpFile)
	value := strings.Split(strings.Split(content, "\n")[7], " ")
	tcp.ActiveOpens, _ = strconv.ParseFloat(value[5], 64)
	tcp.PassiveOpens, _ = strconv.ParseFloat(value[6], 64)
	tcp.AttemptFails, _ = strconv.ParseFloat(value[7], 64)
	tcp.EstabResets, _ = strconv.ParseFloat(value[8], 64)
	tcp.CurrEstab, _ = strconv.ParseFloat(value[9], 64)
	tcp.InSegs, _ = strconv.ParseFloat(value[10], 64)
	tcp.OutSegs, _ = strconv.ParseFloat(value[11], 64)
	tcp.RetransSegs, _ = strconv.ParseFloat(value[12], 64)
	tcp.InErrs, _ = strconv.ParseFloat(value[13], 64)
	return tcp
}

func tcpAvg() module.TcpMetric {
	tcp.ActiveOpens = Delta(tcpNow.ActiveOpens, tcpPre.ActiveOpens)
	tcp.PassiveOpens = Delta(tcpNow.PassiveOpens, tcpPre.PassiveOpens)
	tcp.InSegs = Delta(tcpNow.InSegs, tcpPre.InSegs)
	tcp.OutSegs = Delta(tcpNow.OutSegs, tcpPre.OutSegs)
	tcp.EstabResets = Delta(tcpNow.EstabResets, tcpPre.EstabResets)
	tcp.AttemptFails = Delta(tcpNow.AttemptFails, tcpPre.AttemptFails)
	tcp.CurrEstab = tcpNow.CurrEstab
	tcp.RetransSegs = Delta(tcpNow.RetransSegs, tcpNow.RetransSegs)

	return tcp
}
