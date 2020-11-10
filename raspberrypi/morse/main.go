package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"github.com/jaymzee/morse"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

var pinNumber int

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s pin message\n", os.Args[0])
	}
}

func main() {
	// configure logging
	log.SetFlags(log.Ltime | log.Lmicroseconds)

	// parse flags and check usage
	flag.Parse()
	args := flag.Args()
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}
	var err error
	pinNumber, err = strconv.Atoi(args[0])
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}
	message := strings.ToUpper(strings.Join(args[1:], " "))

	// setup the remaining things
	led, err := gpio0.OpenLED(pinNumber)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	setupCtrlCHandler(led)

	// loop sending message forever
	fmt.Printf("Sending morse code on gpio pin %d\n", pinNumber)
	for {
		morse.Send(led, []byte(message))
	}
}

// CTRL-C handler to turn off the LED
func setupCtrlCHandler(led *gpio0.LED) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	go func() {
		<-c
		led.Off()
		os.Exit(0)
	}()
}
