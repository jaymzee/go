package main

import (
	"flag"
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
)

const (
	// Width of window
	Width   = 1000
	// Height of window
	Height  = 1000
	// CenterX is the center of screen (horizontal)
	CenterX = Width / 2
	// CenterY is the center of screen (vertical)
	CenterY = Height / 2
	// Radius is the radius of the circle of points
	Radius  = 450
	// FPS is frames per second
	FPS     = 1000
	// Points is the total number of points
	Points  = 500
)

// Colors
var (
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 0}
)

// Fonts
var (
	Sans18 *ttf.Font
)

func main() {
	// process program arguments
	rendererFlags := uint32(sdl.RENDERER_ACCELERATED)
	sFlag := flag.Bool("s", false, "use software rendering")
	flag.Parse()
	if *sFlag {
		rendererFlags = uint32(sdl.RENDERER_SOFTWARE)
	}

	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, err := sdl.CreateWindow("Heart",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, rendererFlags)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()
	Sans18, err = ttf.OpenFont("DejaVuSans.ttf", 18)
	if err != nil {
		panic(err)
	}

	// rendering loop
	factor := 1.0
	for running := true; running; {
		DrawFrame(renderer, factor)
		renderer.Present()
		sdl.Delay(uint32(math.Round(1000.0 / FPS)))
		factor += 0.001

		if event := sdl.PollEvent(); event != nil {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}
}

func DrawFrame(renderer *sdl.Renderer, factor float64) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 255)
	var x1, y1, x2, y2 float64
	for n := 0; n < Points; n++ {
		n1 := float64(n)
		n2 := factor * n1
		theta := 2 * math.Pi * n1 / Points
		phi := 2 * math.Pi * n2 / Points
		x1 = Radius * math.Cos(theta)
		y1 = Radius * math.Sin(theta)
		x2 = Radius * math.Cos(phi)
		y2 = Radius * math.Sin(phi)
		renderer.DrawLine(
			int32(CenterX-x1), int32(CenterY-y1),
			int32(CenterX-x2), int32(CenterY-y2),
		)
	}

	factorText := fmt.Sprintf("factor: %6.3f", factor)
	DrawText(renderer, factorText, 10, 10, Sans18, Yellow)
}

// DrawText renders a string to screen coordinates x and y in the
// font and color given.  It is a convenience method that
//   - creates surface
//   - a texture from that surface
//   - renders the texture
func DrawText(renderer *sdl.Renderer, text string, x int32, y int32, font *ttf.Font, color sdl.Color) {
	surface, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		panic(err)
	}
	defer surface.Free()
	texture, err := renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	defer texture.Destroy()
	_, _, w, h, err := texture.Query()
	if err != nil {
		panic(err)
	}
	rect := &sdl.Rect{X: x, Y: y, W: w, H: h}
	renderer.Copy(texture, nil, rect)
}
