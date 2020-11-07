package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"os"
	"time"
)

func main() {
	pinFlag := flag.Int("p", -1, "gpio pin number")
	countFlag := flag.Int("n", 4, "blink count")
	intervalFlag := flag.Int("i", 1000, "blink interval (ms)")
	flag.Parse()
	if *pinFlag < 0 {
		fmt.Fprintf(os.Stderr, "Usage: %s -p pin [options]\n", os.Args[0])
		flag.PrintDefaults()
		os.Exit(1)
	}
	interval := time.Duration(*intervalFlag) * time.Millisecond

	led, err := gpio0.OpenLED(*pinFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i := 0; i < *countFlag; i++ {
		led.On()
		time.Sleep(interval / 2)
		led.Off()
		time.Sleep(interval / 2)
	}
}
