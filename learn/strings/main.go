// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func intToStr(i int) {
	fmt.Printf("unicode: %s\n", string(i))
	fmt.Printf("int:     %s\n", strconv.Itoa(i))
}

func int64ToStr(i int64) {
	fmt.Printf("dec:     %s\n", strconv.FormatInt(i, 10))
	fmt.Printf("hex:     0x%s\n", strconv.FormatInt(i, 16))
	fmt.Printf("oct:     0%s\n", strconv.FormatInt(i, 8))
	fmt.Printf("bin:     %s\n", strconv.FormatInt(i, 2))
}

func strToInt(s string) {
	b := []int{10, 16, 8, 2}
	for _, b := range(b) {
		n, err := strconv.ParseInt(s, b, 64)
		if err == nil {
			fmt.Printf("%s base %2d 0x%04x\n", s, b, n)
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
	s := fmt.Sprintf("%+8d", 97)
	fmt.Println(s)
}
