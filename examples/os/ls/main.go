// ReadDir example, similar to:
//  $ ls -l
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}
	for _, arg := range args {
		ls(arg)
	}
}

func ls(p string) {
	// arg is directory or file?
	info, err := os.Stat(p)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if info.IsDir() {
		// read the directory
		dir, err := os.ReadDir(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		// print the directory
		err = PrintDirEntries(os.Stdout, dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		// it's just a file so print the info for that
		err = PrintFileInfo(os.Stdout, info, 1)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	}
}

func PrintDirEntries(w io.Writer, entries []os.DirEntry) error {
	// find the largest size and get the length of it's string repr
	var maxSize int64 = 0
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		if info.Size() > maxSize {
			maxSize = info.Size()
		}
	}
	sizewidth := len(strconv.FormatInt(maxSize, 10))

	// write the directory to the writer
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return err
		}
		PrintFileInfo(w, info, sizewidth)
	}

	return nil
}

func PrintFileInfo(w io.Writer, info os.FileInfo, sizwidth int) error {
	fmt.Fprintf(w, "%s", info.Mode())
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		fmt.Fprintf(w, " %d %d %d", stat.Nlink, stat.Uid, stat.Gid)
	}
	mtime := info.ModTime().Format("Jan 02 15:04")
	size := padLeft(strconv.FormatInt(info.Size(), 10), ' ', sizwidth)
	fmt.Fprintf(w, " %s %12s %s\n", size, mtime, info.Name())
	return nil
}

func padLeft(str string, pad rune, length int) string {
	if length > len(str) {
		count := length - len(str)
		return strings.Repeat(string(pad), count) + str
	}
	return str
}
