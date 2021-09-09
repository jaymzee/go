package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

// comment with some utf8 café
func main() {
	var (
		fname string
		b     []byte
		err   error
	)

	if len(os.Args) > 1 {
		fname = os.Args[1]
		b, err = os.ReadFile(fname)
		if err != nil {
			panic(err)
		}
	} else {
		fname = "-"
		var buf bytes.Buffer
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}
		b = buf.Bytes()
	}

	r := bytes.NewReader(b)
	lines := count(r, bufio.ScanLines)
	r.Reset(b)
	words := count(r, bufio.ScanWords)
	r.Reset(b)
	runes := count(r, bufio.ScanRunes)
	r.Reset(b)
	bytes := count(r, bufio.ScanBytes)

	fmt.Printf("%3d %3d %3d %3d %s\n", lines, words, runes, bytes, fname)
}

func count(r io.Reader, f bufio.SplitFunc) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(f)
	n := 0
	for scanner.Scan() {
		n++
	}
	return n
}
