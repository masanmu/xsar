package funcs

import (
	"fmt"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"strconv"
	"strings"
	"time"
)

var cpuNow, cpuPre, cpu module.CpuMetric

func CpuMetrics() interface{} {
	cpuPre = cpuMetrics()
	time.Sleep(time.Second)
	cpuNow = cpuMetrics()
	return cpuAvg()
}

func cpuMetrics() module.CpuMetric {
	content := strings.TrimRight(Open(config.CpuFile), "\n")
	value := strings.Split(fmt.Sprintf("%s", strings.Split(content, "\n")[0]), " ")
	cpu.User, _ = strconv.ParseFloat(value[2], 64)
	cpu.Nice, _ = strconv.ParseFloat(value[3], 64)
	cpu.System, _ = strconv.ParseFloat(value[4], 64)
	cpu.Idle, _ = strconv.ParseFloat(value[5], 64)
	cpu.Iowait, _ = strconv.ParseFloat(value[6], 64)
	cpu.Irq, _ = strconv.ParseFloat(value[7], 64)
	cpu.Softirq, _ = strconv.ParseFloat(value[8], 64)
	cpu.Stealstolen, _ = strconv.ParseFloat(value[9], 64)
	cpu.Guest, _ = strconv.ParseFloat(value[10], 64)

	return cpu
}

func cpuAvg() module.CpuMetric {
	totalCpuTimeNow := cpuNow.User + cpuNow.Nice + cpuNow.System + cpuNow.Idle + cpuNow.Iowait + cpuNow.Irq + cpuNow.Softirq + cpuNow.Stealstolen + cpuNow.Guest
	totalCpuTimePre := cpuPre.User + cpuPre.Nice + cpuPre.System + cpuPre.Idle + cpuPre.Iowait + cpuPre.Irq + cpuPre.Softirq + cpuPre.Stealstolen + cpuPre.Guest
	totalCpuTime := totalCpuTimeNow - totalCpuTimePre
	cpu.User = Counter(totalCpuTime, cpuNow.User, cpuPre.User)
	cpu.Nice = Counter(totalCpuTime, cpuNow.Nice, cpuPre.Nice)
	cpu.System = Counter(totalCpuTime, cpuNow.System, cpuPre.System)
	cpu.Idle = Counter(totalCpuTime, cpuNow.Idle, cpuPre.Idle)
	cpu.Iowait = Counter(totalCpuTime, cpuNow.Iowait, cpuPre.Iowait)
	cpu.Irq = Counter(totalCpuTime, cpuNow.Irq, cpuPre.Irq)
	cpu.Softirq = Counter(totalCpuTime, cpuNow.Softirq, cpuPre.Softirq)
	cpu.Stealstolen = Counter(totalCpuTime, cpuNow.Stealstolen, cpuPre.Stealstolen)
	cpu.Guest = Counter(totalCpuTime, cpuNow.Guest, cpuPre.Guest)

	return cpu
}
