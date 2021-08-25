package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	rect := image.Rect(0, 0, 16, 16)
	palette := color.Palette([]color.Color{
		color.RGBA{0, 0, 0, 255},
		color.RGBA{255, 0, 0, 255},
		color.RGBA{0, 255, 0, 255},
	})
	img := image.NewPaletted(rect, palette)
	img.Pix[0] = 1
	img.Pix[1] = 2
	img.Pix[16] = 2
	img.Pix[17] = 1

	f, err := os.Create("foo.png")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	enc := png.Encoder{CompressionLevel: png.BestSpeed}
	err = enc.Encode(f, img)
	if err != nil {
		panic(err)
	}
}
