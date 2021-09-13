package main

import (
	"fmt"
	"io"
	"io/fs"
)

const (
	Red      = "1;31"
	Green    = "1;32"
	Yellow   = "1;33"
	Blue     = "1;34"
	Magenta  = "1;35"
	Cyan     = "1;36"
	GreenRev = "7;32"
)

// FileInfoFormat is the maximum width to pad fields to
type FileInfoFormat struct {
	Stat string
	Mode string
	Size string
	Time string
}

func printDir(w io.Writer, entries []fs.FileInfo) {
	for _, info := range entries {
		fmt.Fprintln(w, colorizeFilename(info))
	}
}

func printDirLong(w io.Writer, entries []fs.FileInfo) error {
	// determine field widths
	var (
		maxMode int
		maxSize int64
	)
	for _, info := range entries {
		mode := info.Mode().String()
		if len(mode) > maxMode {
			maxMode = len(mode)
		}
		if info.Size() > maxSize {
			maxSize = info.Size()
		}
	}

	// format strings are built from the calculated field width
	ifmt := FileInfoFormat{
		Mode: fmt.Sprintf("%%%ds", maxMode),
		Stat: calcStatFormatString(entries),
		Size: fmt.Sprintf(" %%%dd", len(itoa(maxSize))),
		Time: " Jan 02 15:04",
	}

	// write the directory entries to the writer
	for _, info := range entries {
		printFileInfo(w, info, &ifmt)
	}

	return nil
}

func printFileInfo(w io.Writer, info fs.FileInfo, ifmt *FileInfoFormat) {
	// mode might be 10 or 11 characters under Linux depending on the file
	mode := fmt.Sprintf(ifmt.Mode, info.Mode())
	if mode[0] == ' ' && mode[1] >= 'A' && mode[1] <= 'Z' {
		// move special file indicators D (device), L (link), S (socket)
		// all the way over to the left because it looks better in a list
		mode = fmt.Sprintf("%c %s", mode[1], mode[2:])
	}

	// write the thing
	io.WriteString(w, mode)
	printStat(w, info, ifmt.Stat)
	fmt.Fprintf(w, ifmt.Size, info.Size())
	mtime := info.ModTime().Format(ifmt.Time)
	timeFmt := fmt.Sprintf("%%%ds ", len(ifmt.Time))
	fmt.Fprintf(w, timeFmt, mtime)
	fmt.Fprintln(w, colorizeFilename(info))
}

func colorizeFilename(info fs.FileInfo) string {
	name := info.Name()
	if config.cFlag {
		mode := info.Mode()
		perm := int(mode.Perm())
		if mode.IsRegular() {
			if perm&0111 != 0 {
				// if executable
				name = colorText(info.Name(), Green)
			}
		} else if mode&fs.ModeSymlink != 0 {
			name = colorText(info.Name(), Cyan)
		} else if mode&fs.ModeDevice != 0 {
			name = colorText(info.Name(), Yellow)
		} else if mode&fs.ModeNamedPipe != 0 {
			name = colorText(info.Name(), Yellow)
		} else if mode&fs.ModeSocket != 0 {
			name = colorText(info.Name(), Magenta)
		} else if mode&fs.ModeSticky != 0 {
			name = colorText(info.Name(), GreenRev)
		} else if info.IsDir() {
			name = colorText(info.Name(), Blue)
		} else {
			name = colorText(info.Name(), Red)
		}
	}
	return name
}

func colorText(s string, color string) string {
	return fmt.Sprintf("\x1b[%sm%s\x1b[m", color, s)
}
