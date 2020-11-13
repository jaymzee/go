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
	Width = 1000
	// Height of window
	Height = 1000
	// CenterX horizontal center
	CenterX = Width / 2
	// CenterY vertical center
	CenterY = Height / 2
	// Radius is the radius of the circle of points
	Radius = 450
	// FPS is frames per second
	FPS = 1000
	// Points is the total number of points
	Points = 500
)

// Colors
var (
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 128}
)

// Scene holds the state for our scene
type Scene struct {
	Factor float64
	Sans18 *ttf.Font
}

// Init initializes the scene
func (scene *Scene) Init(renderer *sdl.Renderer) {
	scene.Factor = 1.0
	err := renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	if err != nil {
		panic(err)
	}
}

// Draw draws a single frame of the scene
func (scene *Scene) Draw(renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 64)
	var x1, y1, x2, y2 float64
	for n := 0; n < Points; n++ {
		n1 := float64(n)
		n2 := scene.Factor * n1
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

	factor := fmt.Sprintf("factor: %6.3f", scene.Factor)
	DrawText(renderer, factor, 200, 200, scene.Sans18, Yellow)
	scene.Factor += 0.001
}

func main() {
	// process program arguments
	rendererFlags := uint32(sdl.RENDERER_ACCELERATED)
	sFlag := flag.Bool("s", false, "use software rendering")
	flag.Parse()
	if *sFlag {
		rendererFlags = uint32(sdl.RENDERER_SOFTWARE)
	}

	// initialize SDL
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

	// load fonts
	dejaVuSans18, err := ttf.OpenFont("DejaVuSans.ttf", 18)
	if err != nil {
		panic(err)
	}


	// init scene
	scene := Scene{Sans18: dejaVuSans18}
	scene.Init(renderer)
	// scenee render loop
	for running := true; running; {
		scene.Draw(renderer)
		renderer.Present()
		sdl.Delay(uint32(math.Round(1000.0 / FPS)))

		if event := sdl.PollEvent(); event != nil {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
			}
		}
	}
}

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
	err = texture.SetAlphaMod(c.A)
	if err != nil {
		panic(err)
	}
	err = r.Copy(texture, nil, rect)
	if err != nil {
		panic(err)
	}
}
