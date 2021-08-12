package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, _ := os.Stdout.Stat()
	if (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		fmt.Println("terminal")
	} else {
		fmt.Println("not a terminal")
	}
}
