package main

// #cgo CFLAGS:
// #cgo LDFLAGS:
// #include "cursor.h"
import "C"
import (
	"fmt"
	"os"
)

func main() {
	var x, y C.int
	if C.getCursor(&x, &y) != 0 {
		fmt.Fprintf(os.Stderr, "failed to read cursor location")
		os.Exit(1)
	}

	fmt.Printf("%v %v\n", y, x)
}
