package main

import (
	"time"
)

type datetime struct {
	t time.Time
}

func now() (*datetime, error) {
	return &datetime{
		t: time.Now(),
	}, nil
}

func (d *datetime) String() string {
	return d.t.Format("2006-01-02 15:04")
}

func (d *datetime) date() string {
	return d.t.Format("2006-01-02")
}

func (d *datetime) time() string {
	return d.t.Format("15:04")
}
