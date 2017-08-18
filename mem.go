package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type mem struct {
	all    uint64
	free   uint64
	buffer uint64
	cached uint64
}

func memUsage(path string) (*mem, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	// Maps a meminfo metric to its value (i.e. MemTotal --> 100000)
	statMap := make(map[string]uint64)

	for _, line := range lines {
		fields := strings.SplitN(line, ":", 2)
		if len(fields) < 2 {
			continue
		}
		valFields := strings.Fields(fields[1])
		val, _ := strconv.ParseUint(valFields[0], 10, 64)
		statMap[fields[0]] = val
	}
	return &mem{
		all:    statMap["MemTotal"] * KB,
		free:   statMap["MemFree"] * KB,
		buffer: statMap["Buffers"] * KB,
		cached: statMap["Cached"] * KB,
	}, nil
}

func (m *mem) String() string {
	return fmt.Sprintf("%d%% / %dM", 100*(m.all-m.free-m.buffer-m.cached)/m.all, m.all/MB)
}
