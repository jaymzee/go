// +build linux darwin freebsd solaris

package main

import (
	"fmt"
	"io"
	"io/fs"
	"syscall"
)

func calcStatFormatString(entries []fs.FileInfo) string {
	var (
		maxNlink uint64
		maxUid   uint64
		maxGid   uint64
	)
	for _, info := range entries {
		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			// cast everything to uint64 to make compatible with some os's
			maxNlink = max(maxNlink, uint64(stat.Nlink))
			maxUid = max(maxUid, uint64(stat.Uid))
			maxGid = max(maxGid, uint64(stat.Gid))
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
