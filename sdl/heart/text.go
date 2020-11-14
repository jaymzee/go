package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// DrawText renders a string to screen coordinates x and y in the font and
// color given.  It is a convenience function that:
//   - creates a surface from the string
//   - creates a texture from that surface
//   - copies the texture to the renderer
func DrawText(r *sdl.Renderer, s string, x, y int32, f *ttf.Font, c sdl.Color) {
	surface, err := f.RenderUTF8Blended(s, c)
	if err != nil {
		panic(err)
	}
	defer surface.Free()
	texture, err := r.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}
	rect := &sdl.Rect{X: x, Y: y, W: w, H: h}
	if err = texture.SetAlphaMod(c.A); err != nil {
		panic(err)
	}
	if err = r.Copy(texture, nil, rect); err != nil {
		panic(err)
	}
}
