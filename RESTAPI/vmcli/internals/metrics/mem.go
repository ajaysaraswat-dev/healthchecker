package metrics

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type MemoryStat struct {
	Total uint64
	Free  uint64
	Used  uint64
}

func GetMemory() (*MemoryStat, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var total, free uint64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] == "MemTotal:" {
			total, _ = strconv.ParseUint(fields[1], 10, 64)
		}
		if fields[0] == "MemAvailable:" {
			free, _ = strconv.ParseUint(fields[1], 10, 64)
		}
	}
	used := total - free
	return &MemoryStat{
		Total: total,
		Free:  free,
		Used:  used,
	}, nil
}
