package main

import (
	"fmt"
	"github.com/jaymzee/gpio0"
	"os"
	"time"
	"flag"
)

func main() {
	pinFlag := flag.Int("p", -1, "gpio pin number")
	flag.Parse()
	pin := *pinFlag
	if pin < 0 {
		fmt.Fprintln(os.Stderr, "Usage: button -p pin")
		os.Exit(1)
	}

	button, err := gpio0.OpenButton(pin)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for {
		fmt.Printf("%v\n", button.Pressed())
		time.Sleep(250 * time.Millisecond)
	}
}
