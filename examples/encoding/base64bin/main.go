package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	data := []byte("hello world")
	fmt.Printf("%s\n", data)

	enc := base64.StdEncoding

	encoded := make([]byte, enc.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, data)
	fmt.Println("encoded:", encoded)

	decoded := make([]byte, enc.DecodedLen(len(encoded)))
	decodedlen, err := base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		panic(err)
	}
	decoded = decoded[:decodedlen]
	fmt.Printf("decoded: %v\n", decoded)
}
