package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: imginfo image")
		os.Exit(2)
	}

	imgfile, err := os.Open(args[0])
	if err != nil {
		panic(err)
	}
	defer imgfile.Close()

	img, imgfmt, err := image.Decode(imgfile)
	if err != nil {
		panic(err)
	}
	fmt.Println(imgfmt)
	fmt.Printf("%T\n", img)

	fmt.Println(img.ColorModel())
	if img, ok := img.(*image.Paletted); ok {
		fmt.Printf("Paletted")
		for i, c := range img.Palette {
			fmt.Println(i, c)
		}
		fmt.Println(len(img.Pix))
	}
	if img, ok := img.(*image.RGBA); ok {
		fmt.Println(len(img.Pix))
	}
}
