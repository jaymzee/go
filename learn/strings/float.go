// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func floatToStr(x float64) {
	fmt.Printf("%g %q\n", x, strconv.FormatFloat(x, 'g', 4, 64))
	fmt.Printf("%g %q\n", x, strconv.FormatFloat(x, 'E', -1, 64))
	fmt.Printf("%g %q\n", x, strconv.FormatFloat(x, 'E', 4, 64))
}

func strToFloat(s string) {
	x, err := strconv.ParseFloat(s, 64)
	if err == nil {
		fmt.Printf("%q %g\n", s, x)
	}
}

