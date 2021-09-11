package main

import (
	"fmt"
	"io/ioutil"
	"syscall"
)

func main() {
	results, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}

	for _, info := range results {
		mtime := info.ModTime().Format("Jan 2 15:04")
		mode := info.Mode().String()
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			fmt.Printf("%s %d %v %v %7d %s %s\n", mode, stat.Nlink, stat.Uid,
				stat.Gid, info.Size(), mtime, info.Name())
		} else {
			fmt.Printf("%s %7d %d %s\n", mode, info.Size(), mtime, info.Name())
		}
	}
}
