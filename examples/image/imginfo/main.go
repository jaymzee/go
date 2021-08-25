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

	fmt.Printf("%s %T %v\n", imgfmt, img, img.Bounds().Max)

	if img, ok := img.(*image.Paletted); ok {
		fmt.Println("Palette:")
		for i, c := range img.Palette {
			fmt.Println(i, c)
		}
		fmt.Printf("pixels: %v\n", len(img.Pix))
	}
	if img, ok := img.(*image.RGBA); ok {
		fmt.Printf("pixels: %v\n", len(img.Pix))
	}
	if img, ok := img.(*image.YCbCr); ok {
		fmt.Printf("pixels: %v\n", len(img.Y))
	}
}
