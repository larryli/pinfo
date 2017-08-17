package main

import (
	"github.com/davecheney/gpio"
	"time"
)

type button struct {
	pin   gpio.Pin
	timer *time.Timer
}

const (
	GPIO0 = gpio.GPIO0
	GPIO1 = gpio.GPIO1
	GPIO3 = gpio.GPIO3
	GPIO7 = gpio.GPIO7
)

func newButton(number int) (*button, error) {
	pin, err := gpio.OpenPin(number, gpio.ModeInput)
	if err != nil {
		return nil, err
	}
	return &button{
		pin: pin,
	}, nil
}

func (b *button) watch(callback func(holden bool)) error {
	if err := b.pin.BeginWatch(gpio.EdgeBoth, func() {
		if b.timer != nil {
			b.timer.Stop()
			b.timer = nil
		}
		if b.pin.Get() {
			b.timer = time.AfterFunc(2*time.Second, func() {
				callback(true)
			})
			callback(false)
		}
	}); err != nil {
		return err
	}
	return nil
}

func (b *button) close() error {
	if err := b.pin.EndWatch(); err != nil {
		return err
	}
	if err := b.pin.Close(); err != nil {
		return err
	}
	return nil
}
