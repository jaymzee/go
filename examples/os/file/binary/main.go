package main

import (
	"io"
	"os"
	"encoding/binary"
)

func main() {
	f, err := os.OpenFile("foo", os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.Write([]byte("data\n"))
	io.WriteString(f, "le:")
	var num int32 = 0x12345678
	err = binary.Write(f, binary.LittleEndian, num)
	if err != nil {
		panic(err)
	}
	io.WriteString(f, "be:")
	err = binary.Write(f, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
}
