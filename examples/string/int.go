// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func intToStr(i int) {
	fmt.Printf("%4d char:    %q\n", i, i) // %q = 'A', %c = A
	fmt.Printf("%4d unicode: %#U\n", i, i) // %U = U+0041, %#U = U+0041 'A'
	fmt.Printf("%4d string:  %q\n", i, string(i))
	fmt.Printf("%4d itoa:    %q\n", i, strconv.Itoa(i))
}

func int64ToStr(i int64) {
	fmt.Printf("%4d dec:     %q\n", i, strconv.FormatInt(i, 10))
	fmt.Printf("%4d hex:     %q\n", i, strconv.FormatInt(i, 16))
	fmt.Printf("%4d oct:     %q\n", i, strconv.FormatInt(i, 8))
	fmt.Printf("%4d bin:     %q\n", i, strconv.FormatInt(i, 2))
}

func strToInt(s string) {
	n, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("%q atoi %d\n", s, n)
	}
	for _, base := range([]int{10, 16, 8, 2}) {
		n, err := strconv.ParseInt(s, base, 64)
		if err == nil {
			fmt.Printf("%q base %2d %#04x\n", s, base, n)
		}
	}
}

