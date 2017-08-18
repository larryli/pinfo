package main

import (
	"fmt"
	"syscall"
)

type disk struct {
	all  uint64
	free uint64
}

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func diskUsage(path string) (*disk, error) {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs(path, &fs)
	if err != nil {
		return nil, err
	}
	return &disk{
		all:  fs.Blocks * uint64(fs.Bsize),
		free: fs.Bfree * uint64(fs.Bsize),
	}, nil
}

func (d *disk) String() string {
	return fmt.Sprintf("%d%% / %.1fG", 100*(d.all-d.free)/d.all, float64(d.all)/float64(GB))
}
