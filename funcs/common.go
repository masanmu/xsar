package funcs

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Delta(now, pre float64) float64 {
	if now >= pre {
		return now - pre
	} else {
		return 0.0
	}
}

func Counter(denominator, now, pre float64) float64 {
	delta := (now - pre) / denominator
	if delta <= 0 {
		return 0.0
	} else {
		return delta
	}
}

func Open(file string) string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatalf("Failed to open %s file", file)
	}
	defer f.Close()

	buf := make([]byte, 4096)

	n, err := f.Read(buf)
	if err != nil {
		log.Fatalf("Failed to read %s file", file)
	}

	return fmt.Sprintf("%s", buf[:n])
}

func FormatPrint(metrics []string) {
	for _, metric := range metrics {
		fmt.Printf("%s\t", metric)
	}
	fmt.Printf("\n")
}

func LivePrint(name string) (interface{}, error) {
	for _, v := range Mappers {
		if name == v.Name {
			return v.Fs(), nil
		}
	}

	return "", errors.New("Not module")
}
