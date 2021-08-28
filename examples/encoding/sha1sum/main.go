package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"log"
	"flag"
	"bufio"
)

func fromStdin() {
		h := sha1.New()
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintln(h, scanner.Text())
		}
		fmt.Printf("%s  -\n", hex.EncodeToString(h.Sum(nil)))
}

func sha1sumBytes(data []byte) string {
	h := sha1.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func sha1sum(filename string) string {
	h := sha1.New()
	if filename == "-" {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Fprintln(h, scanner.Text())
		}
		return hex.EncodeToString(h.Sum(nil))
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Printf("%s  -\n", sha1sum("-"))
		os.Exit(0)
	}

	for _, filename := range args {
		fmt.Printf("%s  %s\n", sha1sum(filename), filename)
	}
}
