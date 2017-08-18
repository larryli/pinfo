package main

import (
	"errors"
	"net"
)

type ip struct {
	ip net.IP
}

func externalIP() (*ip, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			var i net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				i = v.IP
			case *net.IPAddr:
				i = v.IP
			}
			if i == nil || i.IsLoopback() {
				continue
			}
			i = i.To4()
			if i == nil {
				continue // not an ipv4 address
			}
			return &ip{
				ip: i,
			}, nil
		}
	}
	return nil, errors.New("No network")
}

func (i *ip) String() string {
	return i.ip.String()
}
