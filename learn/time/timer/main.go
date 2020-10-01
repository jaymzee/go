package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	fmt.Println("Timer 1 started")
	<-timer1.C
	fmt.Println("Timer 1 fired")
}
