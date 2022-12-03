package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

// The Xmastree has a total of 25 LEDS
//  - Pin 2 for the LED of the star
//  - Pins 4 to 27 for LEDs of the tree
type xmastree struct {
	starLed  rpio.Pin
	treeLeds []rpio.Pin
}

func newXmastree() xmastree {
	if err := rpio.Open(); err != nil {
		log.Fatal(err)
	}

	x := xmastree{
		starLed:  rpio.Pin(2),
		treeLeds: make([]rpio.Pin, 24),
	}

	x.starLed.Output()
	x.starLed.Low()

	for i := range x.treeLeds {
		x.treeLeds[i] = rpio.Pin(4 + i)
		x.treeLeds[i].Output()
		x.treeLeds[i].Low()
	}

	return x
}

func (x *xmastree) stop() {
	for _, led := range x.treeLeds {
		led.Low()
	}
	x.starLed.Low()
	rpio.Close()
}

func (x *xmastree) run() {
	ctrlc := make(chan os.Signal)
	signal.Notify(ctrlc, os.Interrupt, syscall.SIGTERM)

	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	x.runRandom(stop, &wg)

	select {
	case <-ctrlc:
		signal.Stop(ctrlc)
		stop <- struct{}{}
	}
	wg.Wait()
}

func (x *xmastree) runRandom(stop <-chan struct{}, wg *sync.WaitGroup) {
	var i int
	maxLEDs := len(x.treeLeds) - 1

	x.starLed.High()

	ticker := time.NewTicker(30 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ticker.C:
				i = rand.Intn(maxLEDs)
				x.treeLeds[i].Toggle()
			case <-stop:
				x.stop()
				wg.Done()
				return
			}
		}
	}()
}

func (x *xmastree) runTest() {
	allLeds := append([]rpio.Pin{x.starLed}, x.treeLeds...)
	for _, led := range allLeds {
		fmt.Printf("Testing LED n°%02d\n", led)
		for n := 1; n <= 5; n++ {
			led.High()
			time.Sleep(1 * time.Second)
			led.Low()
		}
	}
}
