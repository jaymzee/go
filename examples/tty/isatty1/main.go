package main

import (
	"fmt"
	"github.com/mattn/go-isatty"
	"os"
)

func main() {
	if isatty.IsTerminal(os.Stdout.Fd()) {
		fmt.Println("terminal")
	} else if isatty.IsCygwinTerminal(os.Stdout.Fd()) {
		fmt.Println("Cygwin/MSYS2 terminal")
	} else {
		fmt.Println("not a terminal")
	}
}
