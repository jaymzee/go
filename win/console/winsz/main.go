package main

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS:
// #include <stdlib.h>
// #include "winsz.h"
import "C"
import (
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	var mode uint32
	h, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
			panic(err)
	}
	windows.GetConsoleMode(h, &mode)
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(h, mode)

	var sz C.struct_WindowSize
	C.GetConsoleWindowSize(&sz)
	fmt.Printf("%v %v %v %v\n", sz.rows, sz.cols, sz.xres, sz.yres)
}
