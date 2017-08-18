package main

import (
	"fmt"
	"syscall"
	"time"
)

type info struct {
	up     time.Duration
	loads  [3]float64
	all    uint64
	free   uint64
	buffer uint64
}

const (
	SCALE = 65536.0 // magic
)

func sysInfo() (*info, error) {
	si := syscall.Sysinfo_t{}
	err := syscall.Sysinfo(&si)
	if err != nil {
		return nil, err
	}
	unit := uint64(si.Unit)
	i := &info{
		up:     time.Duration(si.Uptime) * time.Second,
		all:    uint64(si.Totalram) / unit,
		free:   uint64(si.Freeram) / unit,
		buffer: uint64(si.Bufferram) / unit,
	}
	i.loads[0] = float64(si.Loads[0]) / SCALE
	i.loads[1] = float64(si.Loads[1]) / SCALE
	i.loads[2] = float64(si.Loads[2]) / SCALE
	return i, nil
}

func (i *info) uptime() string {
	return i.up.String()
}

func (i *info) loadavg() string {
	return fmt.Sprintf("%.2f %.2f %.2f", i.loads[0], i.loads[1], i.loads[2])
}

func (i *info) memUsage() string {
	return fmt.Sprintf("%d%% / %dM", 100*(i.all-i.free-i.buffer)/i.all, i.all/MB)
}
