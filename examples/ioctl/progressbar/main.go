package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	ONE_MSEC = 1000 * 1000
	NUM      = 100
)

func main() {
	var bar string

	cols, err := TerminalWidth()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\ndefaulting to %d columns", err, cols)
	}

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

func TerminalWidth() (int, error) {
	sizeobj, err := GetWinsize()
	if err != nil {
		return 80, err
	}
	return int(sizeobj.Col), nil
}
