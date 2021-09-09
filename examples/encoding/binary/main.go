package main

import (
	"fmt"
	"encoding/binary"
)

func main() {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, 0xdeadbeef)
	fmt.Printf("%#v\n", b)
	binary.BigEndian.PutUint32(b, 0xdeadbeef)
	fmt.Printf("%#v\n", b)
	x := binary.BigEndian.Uint32(b)
	fmt.Printf("%#08x\n", x)
}
