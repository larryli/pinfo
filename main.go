package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	if release, err := osRelase("/etc/os-release"); err != nil {
		log.Printf("Unable get OS release: %s\n", err)
	} else {
		log.Printf("OS: %s\n", release)
	}

	if kernel, err := uname(); err != nil {
		log.Printf("Unable get uname: %s\n", err)
	} else {
		log.Printf("Kernel: %s\n", kernel)
	}

	if time, err := now(); err != nil {
		log.Printf("Unable get datetime: %s\n", err)
	} else {
		log.Printf("Date time: %s\n", time)
	}

	if info, err := sysInfo(); err != nil {
		log.Printf("Unable get sys info: %s\n", err)
	} else {
		log.Printf("System load: %s\n", info.loadavg())
		log.Printf("Up time: %s\n", info.uptime())
		// log.Printf("Memory usage: %s\n", info.memUsage())
	}

	if mem, err := memUsage("/proc/meminfo"); err != nil {
		log.Printf("Unable get memory usage: %s\n", err)
	} else {
		log.Printf("Memory usage: %s\n", mem)
	}

	if ip, err := externalIP(); err != nil {
		log.Printf("Unable get ip address: %s\n", err)
	} else {
		log.Printf("IP: %s\n", ip)
	}

	if temp, err := cpuTemp(); err != nil {
		log.Printf("Unable get cpu temp: %s\n", err)
	} else {
		log.Printf("CPU temp: %s\n", temp)
	}

	if disk, err := diskUsage("/"); err != nil {
		log.Printf("Unable get disk usage: %s\n", err)
	} else {
		log.Printf("Usage of /: %s\n", disk)
	}

	button, err := newButton(GPIO7)
	if err != nil {
		log.Printf("Error init button: %s\n", err)
		return
	}

	// clean up on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			log.Println("Closing button and terminating program.")
			button.close()
			os.Exit(0)
		}
	}()

	err = button.watch(func(holden bool) {
		if holden {
			log.Println("Button holden.")
		} else {
			log.Println("Button click.")
		}
	})
	if err != nil {
		log.Printf("Unable to watch button: %s\n", err)
		os.Exit(1)
	}
	log.Printf("Now watching button:\n")

	for {
		time.Sleep(2000 * time.Millisecond)
	}
}
