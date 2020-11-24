// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
)

func main() {
	const i = 0xB6 // constants are just numbers (no type)

	intToStr(i)
	int64ToStr(i)
	strToInt("AB")
	strToInt("42")
	strToInt("0110")
	floatToStr(3.141592654)
	strToFloat("3.1415")
	strToFloat("6.02E23")
	boolToStr(false)
	boolToStr(true)
	strToBool("false")
	strToBool("true")

	// print in hex with spaces
	hello := "Hello World!"
	fmt.Printf("fmt \"%% x\", %q: % x\n", hello, hello)

	// print arrays and slices in different radix
	nums := [5]int{10,11,12,13,3}
	fmt.Printf("fmt \"%%x\", %#v: %02x\n", nums, nums)

	// scan
	s := "1.1 2.7 3.14 4.2"
	var a [4]float64
	_, err := fmt.Sscan(s, &a[0], &a[1], &a[2], &a[3])
	if err != nil {
		panic(err)
	}
	fmt.Printf("fmt.Sscan(%q): %v\n", s, a)
}
