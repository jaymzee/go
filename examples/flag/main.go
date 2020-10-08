// command line flags are easy peasy in go
package main

import (
	"flag"
	"fmt"
)

func main() {
	var m int
	flag.IntVar(&m, "m", 42, "a number like 42")
	nFlag := flag.Int("n", 1234, "a number like 1234")
	fFlag := flag.Float64("f", 3.14, "a number like 3.14")
	bFlag := flag.Bool("b", false, "true or false")
	sFlag := flag.String("s", "apple", "a string")
	flag.Parse()

	fmt.Println("-n", *nFlag)
	fmt.Println("-m", m)
	fmt.Println("-b", *bFlag)
	fmt.Println("-f", *fFlag)
	fmt.Printf("-s %q\n", *sFlag)

	for i, v := range flag.Args() {
		fmt.Printf("arg[%d] = %q\n", i, v)
	}
}
