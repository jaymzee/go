package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/go/morse"
	"github.com/jaymzee/gpio0"
	"log"
	"os"
	"os/signal"
	"strings"
)

func main() {
	pinFlag := flag.Int("p", -1, "gpio pin number")
	flag.Parse()
	if *pinFlag < 0 || flag.NArg() < 1 {
		fmt.Println("Usage: morse -p pin message")
		os.Exit(1)
	}

	log.SetFlags(log.Ltime | log.Lmicroseconds)
	led := gpio0.NewLED(*pinFlag)
	msg := strings.Join(flag.Args(), " ")
	setupCtrlCHandler(led)

	fmt.Printf("Sending morse code on gpio pin %d\n", *pinFlag)
	for {
		morse.Send(led, msg)
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
