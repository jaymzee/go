// demonstrate several ways to convert to/from strings
package main

import (
	"fmt"
	"strconv"
)

func boolToStr(b bool) {
	fmt.Printf("%v %q\n", b, strconv.FormatBool(b))
}

func strToBool(s string) {
	b, err := strconv.ParseBool(s)
	if err == nil {
		fmt.Printf("%q %v\n", s, b)
	}
}

