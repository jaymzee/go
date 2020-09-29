// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func intToStr(i int) {
	fmt.Printf("%4d unicode: %q\n", i, string(i))
	fmt.Printf("%4d int:     %q\n", i, strconv.Itoa(i))
}

func int64ToStr(i int64) {
	fmt.Printf("%4d dec:     %q\n", i, strconv.FormatInt(i, 10))
	fmt.Printf("%4d hex:     %q\n", i, strconv.FormatInt(i, 16))
	fmt.Printf("%4d oct:     %q\n", i, strconv.FormatInt(i, 8))
	fmt.Printf("%4d bin:     %q\n", i, strconv.FormatInt(i, 2))
}

func strToInt(s string) {
	b := []int{10, 16, 8, 2}
	for _, b := range(b) {
		n, err := strconv.ParseInt(s, b, 64)
		if err == nil {
			fmt.Printf("%q base %2d 0x%04x\n", s, b, n)
		}
	}
}

func main() {
	const i = 0xB6 // constants are just numbers (no type)

	intToStr(i)
	int64ToStr(i)
	strToInt("AB")
	strToInt("42")
	strToInt("0110")
	fmt.Printf("%q\n", fmt.Sprintf("%+8d", 97))

	// print in hex with spaces
	fmt.Printf("% x\n", "Hello World!")
	sl := [5]int{10,11,12,13,3}
	fmt.Printf("%T %02x\n", sl, sl)
}
