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

	// sprintf
	fmt.Printf("%q\n", fmt.Sprintf("%+8d", 97))

	// print in hex with spaces
	fmt.Printf("% x\n", "Hello World!")
	sl := [5]int{10,11,12,13,3}
	fmt.Printf("%T %02x\n", sl, sl)
}
