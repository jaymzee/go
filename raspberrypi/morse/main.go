package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	// configure logging
	log.SetFlags(log.Ltime | log.Lmicroseconds)

	// flags
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] message\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "options:\n")
		flag.PrintDefaults()
	}
	gflag := flag.Int("g", -1, "gpio pin number")
	oflag := flag.Int("o", 1, "onboard led number")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	message := strings.ToUpper(strings.Join(args, " "))

	var led *gpio0.LED
	var err error

	if *gflag >= 0 {
		led, err = gpio0.OpenGPIO(*gflag, false)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Sending morse code to gpio pin %d\n", *gflag)
	} else {
		led, err = gpio0.OpenLocal(*oflag)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		fmt.Printf("Sending morse code to onboard led %d\n", *oflag)
	}

	setupCtrlCHandler(led)
	for {
		_, err = io.WriteString(led, message)
		if err != nil {
			panic(err)
		}
	}
}

// CTRL-C handler to turn off the LED
func setupCtrlCHandler(led *gpio0.LED) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		if led.Local() && led.Pin() == 1 {
			led.On() // this is the power good (input) led
		} else {
			led.Off() // everything else should be turned off
		}
		os.Exit(0)
	}()
}
