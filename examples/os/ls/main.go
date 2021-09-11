// ReadDir example, similar to:
//  $ ls -l
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	dirname := "."
	if len(os.Args) > 1 {
		dirname = os.Args[1]
	}

	// read the directory
	dir, err := os.ReadDir(dirname)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// find the largest size and get it's width in characters
	var maxSize int64 = 0
	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		if info.Size() > maxSize {
			maxSize = info.Size()
		}
	}
	sizeWidth := len(strconv.FormatInt(maxSize, 10))

	// print the directory
	for _, entry := range dir {
		info, err := entry.Info()
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s", info.Mode())
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			fmt.Printf(" %d %d %d", stat.Nlink, stat.Uid, stat.Gid)
		}
		mtime := info.ModTime().Format("Jan 02 15:04")
		size := padLeft(strconv.FormatInt(info.Size(), 10), ' ', sizeWidth)
		fmt.Printf(" %s %12s %s\n", size, mtime, entry.Name())
	}
}

func padLeft(str string, pad rune, length int) string {
	return strings.Repeat(string(pad), length-len(str)) + str
}
