package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("hello world")
	fmt.Printf("%s\n", data)

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println(encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", decoded)
}
