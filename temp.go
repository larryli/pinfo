package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type temp struct {
	t float64
}

var tempFiles = []string{
	"/etc/armbianmonitor/datasources/soctemp",
	"/sys/devices/virtual/thermal/thermal_zone0/temp",
	"/sys/devices/virtual/thermal/thermal_zone1/temp",
	"/sys/class/thermal/thermal_zone0/temp",
	"/sys/class/thermal/thermal_zone1/temp",
	"/sys/devices/platform/sunxi-i2c.0/i2c-0/0-0034/temp1_input",
	"/sys/devices/platform/a20-tp-hwmon/temp1_input",
}

func parseTemp(dat []byte) (float64, error) {
	val, err := strconv.ParseUint(strings.TrimSpace(string(dat)), 10, 64)
	if err == nil {
		if val > 1000 {
			return float64(val) / 1000, nil
		}
		return float64(val), nil
	}
	return 0, err
}

func cpuTemp() (*temp, error) {
	for _, file := range tempFiles {
		if dat, err := ioutil.ReadFile(file); err == nil {
			if t, err := parseTemp(dat); err == nil {
				return &temp{
					t: t,
				}, nil
			}
		}
	}
	return nil, errors.New("Cannot get cpu temp")
}

func (t *temp) String() string {
	return fmt.Sprintf("%d°C", int(t.t))
}
