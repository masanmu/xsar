package funcs

import (
	"github.com/toolkits/sys"
	"github.com/xsar/config"
	"github.com/xsar/module"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var df []module.DfMetric
var now module.DfMetric

func DfMetrics() interface{} {
	return dfMetrics()
}

func dfMetrics() []module.DfMetric {
	content := Open(config.MntFile)
	for _, line := range strings.Split(content, "\n") {
		if strings.HasPrefix(line, "/") {
			mnt := strings.Split(line, " ")[1]
			outBlock, err := sys.CmdOutNoLn("df", "-k", mnt)
			if err != nil {
				log.Fatalf("Error shell command df -k %s", mnt)
			}
			outInode, err := sys.CmdOutNoLn("df", "-i", mnt)
			if err != nil {
				log.Fatalf("Error shell command df -i %s", mnt)
			}
			compile, _ := regexp.Compile(" +")
			block := compile.ReplaceAllString(strings.Split(outBlock, "\n")[1], "|")
			inode := compile.ReplaceAllString(strings.Split(outInode, "\n")[1], "|")
			now.Mount = mnt
			now.Btotl, _ = strconv.ParseFloat(strings.Split(block, "|")[1], 64)
			now.Bused, _ = strconv.ParseFloat(strings.Split(block, "|")[2], 64)
			now.Bfree, _ = strconv.ParseFloat(strings.Split(block, "|")[3], 64)
			now.Butil = strings.Split(block, "|")[4]
			now.Itotl, _ = strconv.ParseFloat(strings.Split(inode, "|")[1], 64)
			now.Iused, _ = strconv.ParseFloat(strings.Split(inode, "|")[2], 64)
			now.Ifree, _ = strconv.ParseFloat(strings.Split(inode, "|")[3], 64)
			now.Iutil = strings.Split(inode, "|")[4]
			df = append(df, now)
		}
	}
	return df
}
