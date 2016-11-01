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

const (
	MaxList           = 10
	FormatTimeString  = "%c[%d;%d;%dm%20s%c[0m"
	FormatHeadString  = "%c[%d;%d;%dm%14s%c[0m"
	FormatValueString = "%c[0;40;37m%14s%c[0m"
	ColorTag          = 0x1B
	BackGround        = 44 //40-47
	Prospect          = 30 //30-37
	Flag              = 7  //0-终端默认设置,1-高亮显示,4-使用下划线,5-闪烁,7-反白显示,8-不可见
)
