package main

import (
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
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
