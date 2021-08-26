package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer

	fmt.Fprintf(&buf, "Hello")

	fmt.Printf("%#v\n", buf.Bytes())
	fmt.Println(buf.String())
}
