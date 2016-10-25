package funcs

import (
	"github.com/xsar/config"
	"github.com/xsar/module"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type metric struct {
	disk      string
	rdIos     float64
	rdMerges  float64
	rdSectors float64
	rdTicks   float64
	wrIos     float64
	wrMerges  float64
	wrSectors float64
	wrTicks   float64
	ticks     float64
	aveq      float64
}

var ioNow, ioPre []metric
var io []module.IoMetric

func IoMetrics() interface{} {
	ioPre = ioMetrics()
	time.Sleep(time.Second)
	ioNow = ioMetrics()
	return ioAvg()
}

func ioMetrics() []metric {
	content := Open(config.IoFile)
	compile := regexp.MustCompile("sd[a-z] .*")
	value := compile.FindAllString(content, -1)

	var io []metric
	var temp metric
	for _, preDisk := range value {
		diskStats := strings.Split(preDisk, " ")
		temp.disk = diskStats[0]
		temp.rdIos, _ = strconv.ParseFloat(diskStats[1], 64)
		temp.rdMerges, _ = strconv.ParseFloat(diskStats[2], 64)
		temp.rdSectors, _ = strconv.ParseFloat(diskStats[3], 64)
		temp.rdTicks, _ = strconv.ParseFloat(diskStats[4], 64)
		temp.wrIos, _ = strconv.ParseFloat(diskStats[5], 64)
		temp.wrMerges, _ = strconv.ParseFloat(diskStats[6], 64)
		temp.wrSectors, _ = strconv.ParseFloat(diskStats[7], 64)
		temp.wrTicks, _ = strconv.ParseFloat(diskStats[8], 64)
		temp.ticks, _ = strconv.ParseFloat(diskStats[10], 64)
		temp.aveq, _ = strconv.ParseFloat(diskStats[11], 64)
		io = append(io, temp)
	}
	return io
}

func ioAvg() []module.IoMetric {
	var temp module.IoMetric
	for index, disk := range ioNow {
		temp.Disk = disk.disk
		temp.Rrqms = Delta(disk.rdMerges, ioPre[index].rdMerges)
		temp.Wrqms = Delta(disk.wrMerges, ioPre[index].wrMerges)
		temp.Rs = Delta(disk.rdIos, ioPre[index].rdIos)
		temp.Ws = Delta(disk.wrIos, ioPre[index].wrIos)
		temp.Rkbs = Delta(disk.rdSectors/2, ioPre[index].rdSectors/2)
		temp.Wkbs = Delta(disk.wrSectors/2, ioPre[index].wrSectors/2)
		temp.Avgrqsz = ioAvgRq(disk, index)
		temp.Avgqusz = ioAvgQu(disk, index)
		temp.Await = ioWait(disk, index)
		temp.Svctm = ioSvcTime(disk, index)
		temp.Util = ioUtil(disk, index)
		io = append(io, temp)
	}
	return io
}

func ioAvgRq(disk metric, index int) float64 {
	totalSector := (disk.rdSectors - ioPre[index].rdSectors) + (disk.wrSectors - ioPre[index].wrSectors)
	totalRos := (disk.rdIos - ioPre[index].rdIos) + (disk.wrIos - ioPre[index].wrIos)
	ioavgrq := totalSector / totalRos
	if ioavgrq <= 0 || math.IsNaN(ioavgrq) {
		return 0.0
	} else {
		return ioavgrq
	}
}

func ioAvgQu(disk metric, index int) float64 {
	aveq := disk.aveq - ioPre[index].aveq
	return aveq / 1000
}

func ioWait(disk metric, index int) float64 {
	ioTicks := (disk.rdTicks - ioPre[index].rdTicks) + (disk.wrTicks - ioPre[index].wrTicks)
	ios := (disk.rdIos - ioPre[index].rdIos) + (disk.wrIos - ioPre[index].wrIos)
	iowait := ioTicks / ios
	if iowait <= 0 || math.IsNaN(iowait) {
		return 0.0
	} else {
		return iowait
	}
}

func ioSvcTime(disk metric, index int) float64 {
	totalRos := (disk.rdIos - ioPre[index].rdIos) + (disk.wrIos - ioPre[index].wrIos)
	totalUser := disk.ticks - ioPre[index].ticks
	iosvc := totalUser / totalRos
	if iosvc <= 0 || math.IsNaN(iosvc) {
		return 0.0
	} else {
		return iosvc
	}
}

func ioUtil(disk metric, index int) float64 {
	totalUser := disk.ticks - ioPre[index].ticks
	return totalUser / 10.0
}
