package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	data := []byte{1, 2, 3, 4}

	fmt.Println(hex.EncodeToString(data))
}
