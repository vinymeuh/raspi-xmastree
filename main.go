package main

import (
	"flag"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio"
)

const (
	PIN_STAR  = 2
	LEDS      = 24 // 24 leds from 4 to 27 (nothing on 3 ?)
	FIRST_LED = 4
	PERIOD    = 30 * time.Millisecond
)

type Xmastree struct {
	Star rpio.Pin
	Leds []rpio.Pin
}

func NewXmastree() *Xmastree {
	x := &Xmastree{}

	x.Star = rpio.Pin(PIN_STAR)
	x.Star.Output()
	x.Star.High()

	x.Leds = make([]rpio.Pin, LEDS)
	for i := 0; i < LEDS; i++ {
		pin := rpio.Pin(i + FIRST_LED)
		pin.Output()
		x.Leds[i] = pin
	}
	return x
}

func (x *Xmastree) Reset() {
	for _, pin := range x.Leds {
		pin.Low()
	}
	x.Star.Low()
}

func (x *Xmastree) Run(stop <-chan struct{}, wg *sync.WaitGroup) {
	ticker := time.NewTicker(PERIOD)
	go func() {
		for {
			select {
			case <-ticker.C:
				i := rand.Intn(LEDS)
				x.Leds[i].Toggle()
			case <-stop:
				x.Reset()
				wg.Done()
				return
			}
		}
	}()
}

func main() {
	reset := flag.Bool("reset", false, "turn off all leds")
	flag.Parse()

	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}
	defer rpio.Close()

	x := NewXmastree()

	if *reset == true {
		x.Reset()
		os.Exit(0)
	}

	ctrlc := make(chan os.Signal)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	x.Run(stop, &wg)

	select {
	case <-ctrlc:
		signal.Stop(ctrlc)
		stop <- struct{}{}
	}
	wg.Wait()
}
