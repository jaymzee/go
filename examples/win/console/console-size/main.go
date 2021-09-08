package main

import (
	"fmt"
	"golang.org/x/sys/windows"
)

func main() {
	h, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}

	// enable virtual terminal processing
	var mode uint32
	windows.GetConsoleMode(h, &mode)
	mode |= windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING
	windows.SetConsoleMode(h, mode)

	// save cursor, move bottom right
	// fmt.Printf("\x1b[s\x1b[999;999H")

	// read position from stdin
	/*
		reader := bufio.NewReader(os.Stdin)
		s, err := reader.ReadString('R')
		if err != nil {
			panic(err)
		}
		// parse terminal query result
		re := regexp.MustCompile(`\[(\d+);(\d+)`)
		results := re.FindAllString(s, 2)
		fmt.Println("regex:", results)
	*/

	// get rows and cols from ConsoleScreenBufferInfo
	var csbi windows.ConsoleScreenBufferInfo
	windows.GetConsoleScreenBufferInfo(h, &csbi)
	cols := csbi.Window.Right - csbi.Window.Left + 1
	rows := csbi.Window.Bottom - csbi.Window.Top + 1

	/*
	if csbi.CursorPosition.X > cols {
		rows = csbi.CursorPosition.X
	}
	// doesn't work because it actually reports the buffer lines instead of visible lines
	if csbi.CursorPosition.Y > rows {
		rows = csbi.CursorPosition.Y
	}
	*/

	// restore cursor
	// fmt.Printf("\x1b[u")

	// print results
	fmt.Printf("%d %d\n", rows, cols)
}
