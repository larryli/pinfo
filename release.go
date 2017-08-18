package main

import (
	"io/ioutil"
	"strings"
)

type release struct {
	name    string
	version string
	pretty  string
	code    string
}

func osRelase(path string) (*release, error) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(data), "\n")

	// Maps a meminfo metric to its value (i.e. MemTotal --> 100000)
	statMap := make(map[string]string)

	for _, line := range lines {
		fields := strings.SplitN(line, "=", 2)
		if len(fields) < 2 {
			continue
		}
		statMap[fields[0]] = strings.Trim(strings.TrimSpace(fields[1]), "\"")
	}
	return &release{
		name:    statMap["NAME"],
		version: statMap["VERSION_ID"],
		pretty:  statMap["PRETTY_NAME"],
		code:    statMap["VERSION_CODENAME"],
	}, nil
}

func (r *release) String() string {
	return r.pretty
}
