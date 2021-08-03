package main

import (
	"os"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

const (
	ONE_MSEC    = 1000 * 1000
	_TIOCGWINSZ = 0x5413 // On OSX use 1074295912. Thanks zeebo
	NUM         = 100
)

func main() {
	var bar string

	cols := TerminalWidth()

	for i := 0; i <= NUM; i++ {
		bar = progress(i, NUM, cols)
		os.Stdout.Write([]byte(bar + "\r"))
		os.Stdout.Sync()
		time.Sleep(ONE_MSEC * 50)
	}
	os.Stdout.Write([]byte("\n"))
}

func Bold(str string) string {
	return "\033[1m" + str + "\033[0m"
}

func Highlight(str string) string {
	return "\033[1;32m" + str + "\033[0m"
}

func progress(current, total, cols int) string {
	prefix := strconv.Itoa(current) + " / " + strconv.Itoa(total)
	bar_start := " ["
	bar_end := "] "

	bar_size := cols - len(prefix+bar_start+bar_end)
	amount := int(float32(current) / (float32(total) / float32(bar_size)))
	remain := bar_size - amount

	bar := strings.Repeat("█", amount) + strings.Repeat(" ", remain)

	return Bold(prefix) + bar_start + Highlight(bar) + bar_end
}

func TerminalWidth() int {
	sizeobj, err := GetWinsize()
	if err != nil {
		return 80
	}
	return int(sizeobj.Col)
}

type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16
	Ypixel uint16
}

func GetWinsize() (*Winsize, error) {
	ws := new(Winsize)

	r1, _, errno := syscall.Syscall(syscall.SYS_IOCTL,
		uintptr(syscall.Stdin),
		uintptr(_TIOCGWINSZ),
		uintptr(unsafe.Pointer(ws)),
	)

	if int(r1) == -1 {
		return nil, os.NewSyscallError("GetWinsize", errno)
	}
	return ws, nil
}
