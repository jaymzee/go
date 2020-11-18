package main

import (
	"fmt"
	"time"
)

func main() {
	for t := range time.Tick(500 * time.Millisecond) {
		fmt.Println(t)
	}
}
