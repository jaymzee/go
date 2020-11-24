// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func intToStr(i int) {
	fmt.Printf("fmt \"%%c\",  %d: %c\n", i, i) // %q = 'A', %c = A
	fmt.Printf("fmt \"%%#U\", %d: %#U\n", i, i) // %U = U+0041, %#U = U+0041 'A'
	fmt.Printf("string(%d): %q\n", i, string(i))
	fmt.Printf("strconv.Itoa(%d): %q\n", i, strconv.Itoa(i))
}

func int64ToStr(i int64) {
	fmt.Printf("strconv.FormatInt(%d, 10): %q\n", i, strconv.FormatInt(i, 10))
	fmt.Printf("strconv.FormatInt(%d, 16): %q\n", i, strconv.FormatInt(i, 16))
	fmt.Printf("strconv.FormatInt(%d,  8): %q\n", i, strconv.FormatInt(i, 8))
	fmt.Printf("strconv.FormatInt(%d,  2): %q\n", i, strconv.FormatInt(i, 2))
}

func strToInt(s string) {
	n, err := strconv.Atoi(s)
	if err == nil {
		fmt.Printf("strconv.Atoi(%q): %d\n", s, n)
	}
	for _, base := range([]int{10, 16, 8, 2}) {
		n, err := strconv.ParseInt(s, base, 64)
		if err == nil {
			fmt.Printf("strconv.ParseInt(%q, %2d): %#04x\n", s, base, n)
		}
	}
}

