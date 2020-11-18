package main

import (
	"fmt"
	"time"
)

func main() {
	ch := time.Tick(500 * time.Millisecond)

	fmt.Println("hello")
	<-ch
	fmt.Println(3)
	<-ch
	fmt.Println(2)
	<-ch
	fmt.Println(1)
	for i := 0; i < 3; i++ {
		t := <-ch
		fmt.Printf("%d, %v\n", i, t)
	}
}
