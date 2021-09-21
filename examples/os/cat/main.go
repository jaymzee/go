package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		io.Copy(os.Stdout, os.Stdin)
		os.Exit(0)
	}

	for _, arg := range flag.Args() {
		err := cat(arg)
		if err != nil {
			fmt.Fprintln(os.Stdout, err)
			os.Exit(1)
		}
	}
}

func cat(filename string) error {
	r, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer r.Close()

	io.Copy(os.Stdout, r)

	return nil
}
