// ReadDir example, similar to:
//  $ ls -l
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"syscall"
)

// InfoFormat is the maximum width to pad fields to
type InfoFormat struct {
	Stat  string
	Mode  string
	Size  string
	Time  string
}

func main() {
	args := []string{"."}
	if len(os.Args) > 1 {
		args = os.Args[1:]
	}

	var dir []os.FileInfo
	for _, arg := range args {
		d, err := readDir(arg)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		dir = append(dir, d...)
	}

	err := printDir(os.Stdout, dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func readDir(p string) ([]os.FileInfo, error) {
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
		infos := make([]os.FileInfo, len(dir))
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
	return []os.FileInfo{info}, nil
}

func printDir(w io.Writer, entries []os.FileInfo) error {
	// find the largest width of fields
	var (
		maxMode  int
		maxSize  int64
		maxNlink uint32
		maxUid   uint32
		maxGid   uint32
	)
	for _, info := range entries {
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			if stat.Nlink > maxNlink {
				maxNlink = stat.Nlink
			}
			if stat.Uid > maxUid {
				maxUid = stat.Uid
			}
			if stat.Gid > maxGid {
				maxGid = stat.Gid
			}
		}
		mode := info.Mode().String()
		if len(mode) > maxMode {
			maxMode = len(mode)
		}
		if info.Size() > maxSize {
			maxSize = info.Size()
		}
	}
	// format strings dynamically set base on maximum field width
	ifmt := InfoFormat{
		Mode: fmt.Sprintf("%%%ds", maxMode),
		Stat: fmt.Sprintf(" %%%dd %%%dd %%%dd",
			len(utoa(maxNlink)), len(utoa(maxUid)), len(utoa(maxGid))),
		Size: fmt.Sprintf(" %%%dd", len(itoa(maxSize))),
		Time: " Jan 02 15:04",
	}
	// write the directory entries to the writer
	for _, info := range entries {
		printFileInfo(w, info, &ifmt)
	}

	return nil
}

func printFileInfo(w io.Writer, info os.FileInfo, ifmt *InfoFormat) error {
	// mode might be 10 or 11 characters under Linux depending on the file
	mode := fmt.Sprintf(ifmt.Mode, info.Mode())
	if mode[0] == ' ' && mode[1] >= 'A' && mode[1] <= 'Z' {
		// move special file indicators D (device), L (link), S (socket)
		// all the way over to the left because it looks better in a list
		mode = fmt.Sprintf("%c %s", mode[1], mode[2:])
	}
	io.WriteString(w, mode)
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		fmt.Fprintf(w, ifmt.Stat, stat.Nlink, stat.Uid, stat.Gid)
	}
	fmt.Fprintf(w, ifmt.Size, info.Size())
	mtime := info.ModTime().Format(ifmt.Time)
	timeFmt := fmt.Sprintf("%%%ds", len(ifmt.Time))
	fmt.Fprintf(w, timeFmt, mtime)
	fmt.Fprintf(w, " %s\n", info.Name())

	return nil
}

func utoa(i uint32) string {
	return strconv.FormatUint(uint64(i), 10)
}

func itoa(i int64) string {
	return strconv.FormatInt(i, 10)
}
