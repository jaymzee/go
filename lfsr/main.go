package main

import (
	"fmt"
)

func main() {
	const poly = 0xaa2255dd
	const iv = 0x1

	period := 0
	x := iv
	for {
		fb := x & 1
		x = x >> 1
		if fb == 1 {
			x = x ^ poly
		}
		period++
		if x == iv {
			break
		}
	}
	fmt.Println(period)
}
