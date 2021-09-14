// +build windows plan9

package main

import (
	"io"
	"io/fs"
)

func calcStatFormatString(entries []fs.FileInfo) string {
	// doesn't matter since printStat will be a noop
	return " %d %d %d"
}

func printStat(w io.Writer, info fs.FileInfo, fmtstr string) {
}
