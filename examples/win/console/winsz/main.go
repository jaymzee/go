package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS:
// #include <stdlib.h>
// #include "winsz.h"
import "C"
import "fmt"

func main() {
	var sz C.struct_WindowSize
	C.GetConsoleWindowSize(&sz)
	fmt.Printf("%v %v %v %v\n", sz.rows, sz.cols, sz.xres, sz.yres)
}
