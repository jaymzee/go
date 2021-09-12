// +build linux darwin

package main

import (
	"fmt"
	"io"
	"io/fs"
	"syscall"
)

func calcStatFormatString(entries []fs.FileInfo) string {
	var (
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
	}
	return fmt.Sprintf(" %%%dd %%%dd %%%dd",
		len(utoa(maxNlink)), len(utoa(maxUid)), len(utoa(maxGid)))
}

func printStat(w io.Writer, info fs.FileInfo, fmtstr string) {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		fmt.Fprintf(w, fmtstr, stat.Nlink, stat.Uid, stat.Gid)
	}
}
