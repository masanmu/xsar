package config

const (
	CpuFile     = "/proc/stat"
	IoFile      = "/proc/diskstats"
	LoadFile    = "/proc/loadavg"
	MemFile     = "/proc/meminfo"
	TcpFile     = "/proc/net/snmp"
	UdpFile     = "/proc/net/snmp"
	TrafficFile = "/proc/net/dev"
	MntFile     = "/etc/mtab"
	XsarFile    = "/var/log/xsar.data"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424
	YiB // 1208925819614629174706176
)
