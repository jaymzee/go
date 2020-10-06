package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("sleeping")
	time.Sleep(2 * time.Second)
	fmt.Println("done")
}
