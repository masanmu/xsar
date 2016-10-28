package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var trafficNow, trafficPre, traffic module.TrafficMetric

func TrafficMetrics() interface{} {
	trafficPre = trafficMetrics()
	time.Sleep(time.Second)
	trafficNow = trafficMetrics()
	return trafficAvg()
}

func trafficMetrics() module.TrafficMetric {
	var now module.TrafficMetric
	content := Open(config.TrafficFile)
	values := strings.Split(content, "\n")
	for _, value := range values {
		if strings.HasPrefix(value, "eth") || strings.HasPrefix(value, "em") || strings.HasPrefix(value, "en") {
			trafficInit(strings.Split(value, ":")[1], now)
		}
	}
	return traffic
}

func trafficInit(value string, now module.TrafficMetric) {
	compile, err := regexp.Compile(" +")
	if err != nil {
		log.Fatalf("Failed to initialize regular expression")
	}

	itf := compile.ReplaceAllString(value, "|")
	ifNet := strings.Split(itf, "|")

	byteIn, _ := strconv.ParseFloat(ifNet[1], 64)
	byteOut, _ := strconv.ParseFloat(ifNet[9], 64)
	pkgIn, _ := strconv.ParseFloat(ifNet[2], 64)
	pkgOut, _ := strconv.ParseFloat(ifNet[10], 64)
	pkgDrpIn, _ := strconv.ParseFloat(ifNet[4], 64)
	pkgDrpOut, _ := strconv.ParseFloat(ifNet[12], 64)
	pkgErrIn, _ := strconv.ParseFloat(ifNet[3], 64)
	pkgErrOut, _ := strconv.ParseFloat(ifNet[11], 64)

	now.ByteIn += byteIn
	now.ByteOut += byteOut
	now.PktIn += pkgIn
	now.PktOut += pkgOut
	now.PktDrp += pkgDrpIn + pkgDrpOut
	now.PktErr += pkgErrIn + pkgErrOut
}

func trafficAvg() module.TrafficMetric {
	traffic.ByteIn = Delta(trafficNow.ByteIn, trafficPre.ByteIn)
	traffic.ByteOut = Delta(trafficNow.ByteOut, trafficPre.ByteOut)
	traffic.PktIn = Delta(trafficNow.PktIn, trafficPre.PktIn)
	traffic.PktOut = Delta(trafficNow.PktOut, trafficPre.PktOut)
	traffic.PktErr = Delta(trafficNow.PktErr, trafficPre.PktErr)
	traffic.PktDrp = Delta(trafficNow.PktDrp, trafficPre.PktDrp)
	return traffic
}
