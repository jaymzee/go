package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines) // or ScanWords, ScanRunes, ScanBytes
	dump(scanner)
}

func dump(scanner *bufio.Scanner) {
	n := 1
	for scanner.Scan() {
		fmt.Println(n, scanner.Text())
		n++
	}
}
