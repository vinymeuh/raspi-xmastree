package main

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/stianeikeland/go-rpio"
)

func main() {
	pin := flag.Uint("pin", 0, "pin number [2, 27]")
	flag.Parse()

	if *pin < 2 || *pin > 27 {
		flag.Usage()
		os.Exit(1)
	}

	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	led := rpio.Pin(*pin)
	led.Output()
	led.High()

	ctrlc := make(chan os.Signal)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctrlc:
		led.Low()
	}

}
