package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS:
// #include <stdlib.h>
// #include "fontsize.h"
import "C"
import "fmt"

func main() {
	var sz C.struct_fontsize
	C.get_console_fontsize(&sz)
	fmt.Printf("%v x %v\n", sz.width, sz.height)
}
