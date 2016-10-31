package module

import ()

type Metrics struct {
	Now     int64
	Cpu     CpuMetric
	Df      []DfMetric
	Io      []IoMetric
	Load    LoadMetric
	Mem     MemMetric
	Tcp     TcpMetric
	Udp     UdpMetric
	Traffic TrafficMetric
}

type CpuMetric struct {
	User        float64 `json:"user"`
	Nice        float64 `json:"nice"`
	System      float64 `json:"system"`
	Idle        float64 `json:"idle"`
	Iowait      float64 `json:"iowait"`
	Irq         float64 `json:"irq"`
	Softirq     float64 `json:"softirq"`
	Stealstolen float64 `json:"stealstolen"`
	Guest       float64 `json:"guest"`
}

type DfMetric struct {
	Mount string  `json:"mount"`
	Bfree float64 `json:"bfree"`
	Bused float64 `json:"bused"`
	Btotl float64 `json:"btotl"`
	Butil string  `json:"butil"`
	Ifree float64 `json:"ifree"`
	Iused float64 `json:"iused"`
	Itotl float64 `json:"itotl"`
	Iutil string  `json:"iutil"`
}

type IoMetric struct {
	Disk    string  `json:"zdisk"`
	Rrqms   float64 `json:"rrqms"`
	Wrqms   float64 `json:"wrqms"`
	Rs      float64 `json:"rs"`
	Ws      float64 `json:"ws"`
	Rkbs    float64 `json:"rkbs"`
	Wkbs    float64 `json:"wkbs"`
	Avgrqsz float64 `json:"avgrq-sz"`
	Avgqusz float64 `json:"avgqu-sz"`
	Await   float64 `json:"await"`
	Svctm   float64 `json:"svctm"`
	Util    float64 `json:"%util"`
}

type LoadMetric struct {
	Load1min   float64 `json:"load1min"`
	Load5min   float64 `json:"load5min"`
	Load15min  float64 `json:"load15min"`
	Lastaskpid string  `json:"lastaskpid"`
}

type MemMetric struct {
	MemTotal  float64 `json:"memtotal"`
	MemFree   float64 `json:"memfree"`
	MemUsed   float64 `json:"memused"`
	Cached    float64 `json:"cached"`
	Buffers   float64 `json:"buffers"`
	SwapTotal float64 `json:"swaptotal"`
	SwapFree  float64 `json:"swapfree"`
	SwapUsed  float64 `json:"swapused"`
}

type TcpMetric struct {
	ActiveOpens  float64 `json:"active"`
	PassiveOpens float64 `json:"passive"`
	AttemptFails float64 `json:"attemptfail"`
	EstabResets  float64 `json:"estabresets"`
	CurrEstab    float64 `json:"currestab"`
	InSegs       float64 `json:"insegs"`
	OutSegs      float64 `json:"outsegs"`
	RetransSegs  float64 `json:"retranssegs"`
	InErrs       float64 `json:"InErrs"`
}

type TrafficMetric struct {
	ByteIn  float64 `json:"bytein"`
	ByteOut float64 `json:"byteout"`
	PktIn   float64 `json:"pktin"`
	PktOut  float64 `json:"pktout"`
	PktDrp  float64 `json:"pktdrp"`
	PktErr  float64 `json:"pkgerr"`
}

type UdpMetric struct {
	InDatagrams  float64 `json:"indatagm"`
	NoPorts      float64 `json:"noports"`
	InErrors     float64 `json:"inerrors"`
	OutDatagrams float64 `json:"outdatagm"`
}
