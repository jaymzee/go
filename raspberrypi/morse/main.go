package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"github.com/jaymzee/morse"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmicroseconds)

	// parse flags and check usage
	pinFlag := flag.Int("p", -1, "gpio pin number")
	flag.Parse()
	pin := *pinFlag
	if pin < 0 || flag.NArg() < 1 {
		fmt.Fprintf(os.Stderr, "Usage: %s -p pin message\n", os.Args[0])
		os.Exit(1)
	}
	message := strings.ToUpper(strings.Join(flag.Args(), " "))

	led, err := gpio0.OpenLED(pin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	setupCtrlCHandler(led)

	fmt.Printf("Sending morse code on gpio pin %d\n", pin)
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
