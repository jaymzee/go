package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
)

const (
	// ScreenWidth of window
	ScreenWidth = 1000
	// ScreenHeight of window
	ScreenHeight = 1000
	// FPS is frames per second
	FPS = 1000
)

// Colors
var (
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 128}
)

// Scene contains the state for the scene
type Scene struct {
	Points int
	Factor float64
	Radius float64
	Sans18 *ttf.Font
}

// Init initializes the scene
func (scene *Scene) Init(window *sdl.Window, renderer *sdl.Renderer) {
	var err error
	scene.Factor = 1.0
	scene.Points = 500
	scene.Radius = 450
	scene.Sans18, err = ttf.OpenFont("DejaVuSans.ttf", 18)
	if err != nil {
		panic(err)
	}
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	window.SetTitle("Scene")
}

// Draw draws a single frame of the scene
func (scene *Scene) Draw(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 64)
	var x1, y1, x2, y2 float64
	cx, cy := float64(ScreenWidth/2), float64(ScreenHeight/2)
	for n := 0; n < scene.Points; n++ {
		n1 := float64(n)
		n2 := scene.Factor * n1
		theta := 2 * math.Pi * n1 / float64(scene.Points)
		phi := 2 * math.Pi * n2 / float64(scene.Points)
		x1 = scene.Radius * math.Cos(theta)
		y1 = scene.Radius * math.Sin(theta)
		x2 = scene.Radius * math.Cos(phi)
		y2 = scene.Radius * math.Sin(phi)
		renderer.DrawLine(
			int32(cx-x1), int32(cy-y1),
			int32(cx-x2), int32(cy-y2),
		)
	}

	factor := fmt.Sprintf("factor: %6.3f", scene.Factor)
	DrawText(renderer, factor, 200, 200, scene.Sans18, Yellow)
	scene.Factor += 0.001
}

// Loop is the event loop for the scene
func (scene *Scene) Loop(window *sdl.Window, renderer *sdl.Renderer) {
	scene.Init(window, renderer)
	for running := true; running; {
		scene.Draw(window, renderer)
		renderer.Present()
		sdl.Delay(uint32(math.Round(1000.0 / FPS)))

		if event := sdl.PollEvent(); event != nil {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
			}
		}
	}
}
