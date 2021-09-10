package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		filename := os.Args[1]
		b, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			os.Exit(1)
		}
		wc(b, bufio.ScanLines, bufio.ScanWords, bufio.ScanBytes)
		fmt.Println(filename)
	} else {
		b, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(err)
		}
		wc(b, bufio.ScanLines, bufio.ScanWords, bufio.ScanBytes)
		fmt.Println()
	}
}

func wc(b []byte, splits ...bufio.SplitFunc) {
	r := bytes.NewReader(b)
	for _, f := range splits {
		n := count(r, f)
		r.Reset(b)
		fmt.Printf("%3d ", n)
	}
}

func count(r io.Reader, f bufio.SplitFunc) int {
	var n int
	scanner := bufio.NewScanner(r)
	scanner.Split(f)
	for n = 0; scanner.Scan(); n++ {
	}
	return n
}
