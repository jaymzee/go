package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Printf("%T\n", err)
		panic(err)
	}
	defer file.Close()

	var fileinfo os.FileInfo
	fileinfo, err = file.Stat()
	if err != nil {
		panic(err)
	}

	filesize := fileinfo.Size()
	buffer := make([]byte, filesize)

	_, err = file.Read(buffer)
	if err != nil {
		panic(err)
	}
	
	fmt.Print(string(buffer))
	fmt.Println(buffer)
}
