package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [-reset|-test]\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	reset := flag.Bool("reset", false, "turn off all leds")
	test := flag.Bool("test", false, "test all leds one by one")
	flag.Parse()
	if *reset == true && *test == true {
		flag.Usage()
		os.Exit(1)
	}

	x := newXmastree()

	if *reset == true {
		os.Exit(0)
	}

	if *test == true {
		x.runTest()
		os.Exit(0)
	}

	x.run()
}
