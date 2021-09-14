package main

import (
	"fmt"
	"io"
	"os"
	"flag"
)

func main() {
	flag.Bool("C", true, "Canonical hex+ASCII display")
	flag.Parse()
	args := flag.Args()

	file := os.Stdin // default to reading from stdin
	// if a filename is provided, open that instead of stdin
	if len(args) > 0 {
		f, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}
		defer f.Close()
		file = f
	}

	var addr int
	var buf [16]byte
	// do while there is more stuff to read
	for {
		// get up to 16 bytes
		n, rderr := io.ReadFull(file, buf[:])
		// print the bytes we did get
		printCanonical(addr, buf[:n])
		addr += n

		if rderr == io.EOF || rderr == io.ErrUnexpectedEOF {
			// expecting one of these at some point
			break
		} else if rderr != nil {
			// this is unexpected
			fmt.Fprintln(os.Stdout, rderr)
			os.Exit(1)
		}
	}
}

// canonical hex+ASCII display
func printCanonical(addr int, bytes []byte) {
	i := printHex(addr, bytes)
	if len(bytes) > 0 {
		for ; i < 60; i++ {
			fmt.Print(" ")
		}
		printASCII(bytes)
	}
	fmt.Println()
}

// print bytes as hex, len(bytes) should be between 0 and 16 inclusive
func printHex(addr int, bytes []byte) int {
	fmt.Printf("%08x", addr)
	col := 8
	for i, b := range bytes {
		if i % 8 == 0 {
			fmt.Printf("  %02x", b)
			col += 4
		} else {
			fmt.Printf(" %02x", b)
			col += 3
		}
	}
	return col
}

// print bytes as ASCII if they are printable, len should be 0 to 16 bytes
func printASCII(bytes []byte) {
	if len(bytes) > 0 {
		fmt.Print("|")
		for _, b := range bytes {
			if b > 31 && b < 127 {
				fmt.Printf("%c", b)
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("|")
	}
}
