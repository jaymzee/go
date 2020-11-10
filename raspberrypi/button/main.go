package main

import (
	"flag"
	"fmt"
	"github.com/jaymzee/gpio0"
	"os"
	"time"
	"strconv"
)

var pinNumber int

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s pin\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	// parse program args
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

	button, err := gpio0.OpenButton(pinNumber)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for {
		fmt.Printf("%v\n", button.Pressed())
		time.Sleep(250 * time.Millisecond)
	}
}
