package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
)

var config struct {
	cFlag bool
	lFlag bool
}

func main() {
	EnableANSI()
	flag.BoolVar(&config.cFlag, "color", false, "ansi color")
	flag.BoolVar(&config.lFlag, "l", false, "long format")
	flag.Parse()
	paths := []string{"."}
	if len(flag.Args()) > 0 {
		paths = flag.Args()
	}

	// get directory for each arg and combine into one slice
	var dir []fs.FileInfo
	for _, p := range paths {
		d, err := readDir(p)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		dir = append(dir, d...)
	}

	// print directory to stdout
	if config.lFlag {
		err := printDirLong(os.Stdout, dir)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	} else {
		printDir(os.Stdout, dir)
	}
}

func readDir(p string) ([]fs.FileInfo, error) {
	// p is directory or file?
	info, err := os.Stat(p)
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		// read the directory
		dir, err := os.ReadDir(p)
		if err != nil {
			return nil, err
		}
		infos := make([]fs.FileInfo, len(dir))
		for n, entry := range dir {
			info, err := entry.Info()
			if err != nil {
				return nil, err
			}
			infos[n] = info
		}
		return infos, nil
	}
	// it's just a file and already have the FileInfo
	return []fs.FileInfo{info}, nil
}
