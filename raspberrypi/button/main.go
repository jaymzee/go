package main

import (
	"fmt"
	"github.com/jaymzee/gpio0"
	"time"
)

func main() {
	btn := gpio0.NewButton(19)
	for {
		fmt.Printf("%v\n", btn.Pressed())
		time.Sleep(250 * time.Millisecond)
	}
}
