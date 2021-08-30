// +build linux darwin

package main

import (
	"golang.org/x/sys/unix"
	"os"
)

func GetWinsize() (*Winsize, error) {
	ws, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, os.NewSyscallError("GetWinsize", err)
	}

	w := Winsize(*ws)
	return &w, nil
}
