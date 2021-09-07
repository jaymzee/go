package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"os"
	"strconv"
	"time"
)

var (
	pinNumber     int
	blinkCount    int
	blinkInterval time.Duration
)

func init() {
	flag.IntVar(&blinkCount, "n", 4, "blink count")
	intervalFlag := flag.Int("i", 1000, "blink interval (ms)")
	blinkInterval = time.Duration(*intervalFlag) * time.Millisecond
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s pin [options]\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
}

func main() {
	// parse program arguments
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}
	var err error
	pinNumber, err = strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	led, err := gpio0.OpenGPIO(pinNumber, false)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i := 0; i < blinkCount; i++ {
		led.On()
		time.Sleep(blinkInterval / 2)
		led.Off()
		time.Sleep(blinkInterval / 2)
	}
}
