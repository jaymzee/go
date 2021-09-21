package main

import (
	"flag"
	"io"
	"log"
	"os"
	"fmt"
)

func main() {
	var nFlag bool

	flag.BoolVar(&nFlag, "n", false, "display total bytes copied")
	flag.Parse()

	var total int64 = 0

	if len(flag.Args()) == 0 {
		n, err := io.Copy(os.Stdout, os.Stdin)
		if err != nil {
			log.Fatal(err)
		}
		total = n
	}
	for _, arg := range flag.Args() {
		n, err := cat(arg)
		if err != nil {
			log.Fatal(err)
		}
		total += n
	}

	if nFlag {
		fmt.Fprintf(os.Stderr, "%d\n", total)
	}
}

func cat(filename string) (int64, error) {
	r, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	defer r.Close()

	n, err := io.Copy(os.Stdout, r)
	if err != nil {
		return 0, err
	}

	return n, nil
}
