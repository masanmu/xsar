package xsar

import (
	"os"
)

func Cron() {
	cronFile := "/etc/cron.d/xsar"
	f, err := os.Create(cronFile)
	defer f.Close()
	if err != nil {
		os.Exit(-1)
	}

	f.WriteString("* * * * * root /home/masen/workspace/src/github.com/xsar/main -cron > /dev/null 2>&1")
}
