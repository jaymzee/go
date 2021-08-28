package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"os/exec"
)

func main() {
	var csbi windows.ConsoleScreenBufferInfo
	h, err := windows.GetStdHandle(windows.STD_OUTPUT_HANDLE)
	if err != nil {
		panic(err)
	}
	windows.GetConsoleScreenBufferInfo(h, &csbi)
	cols := csbi.Window.Right - csbi.Window.Left + 1
	rows := csbi.Window.Bottom - csbi.Window.Top + 1

	out, err := exec.Command("stty", "size").Output()
	if err == nil {
		fmt.Sscanf(string(out), "%d %d", &rows, &cols)
	}

	fmt.Printf("%d %d\n", rows, cols)
}
