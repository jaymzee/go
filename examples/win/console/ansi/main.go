package main

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

	fmt.Printf("\x1b[31mred\x1b[m apple\n")
	fmt.Printf("\x1b[32mgreen\x1b[m apple\n")
}
