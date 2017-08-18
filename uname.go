package main

import (
	"syscall"
)

type unames struct {
	os      string
	node    string
	release string
	version string
	arch    string
}

func uname() (*unames, error) {
	utsname := syscall.Utsname{}
	if err := syscall.Uname(&utsname); err != nil {
		return nil, err
	}
	return &unames{
		os:      byte2string(utsname.Sysname),
		node:    byte2string(utsname.Nodename),
		release: byte2string(utsname.Release),
		version: byte2string(utsname.Version),
		arch:    byte2string(utsname.Machine),
	}, nil
}

func byte2string(bs [65]uint8) string {
	b := make([]byte, len(bs))
	for i, v := range bs {
		b[i] = byte(v)
	}
	return string(b)
}

func (u *unames) String() string {
	return u.release
}
