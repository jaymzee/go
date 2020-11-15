package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/gfx"
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
	FPS = 200
)

// Colors
var (
	Red  = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	Green  = sdl.Color{R: 0, G: 255, B: 0, A: 64}
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
)

// Scene contains the state for the scene
type Scene struct {
	points int
	factor float64
	radius float64
	sans18 *ttf.Font
}

// Loop is the event loop for the scene
func (scene *Scene) Loop(window *sdl.Window, renderer *sdl.Renderer) {
	scene.init(window, renderer)
	fpsmgr := new(gfx.FPSmanager)
	gfx.InitFramerate(fpsmgr)
	gfx.SetFramerate(fpsmgr, FPS)
	for running := true; running; {
		// respond to events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				fmt.Println("Quit")
				running = false
			default:
				fmt.Printf("%T %#v\n", event, event)
			}
		}

		// draw a single frame of the scene
		scene.draw(window, renderer)
		renderer.Present()
		gfx.FramerateDelay(fpsmgr)

		// update scene
		scene.factor += 0.001
	}
}

// NewScene initializes the scene
func (scene *Scene) init(window *sdl.Window, renderer *sdl.Renderer) {
	scene.factor = 1.0
	scene.points = 500
	scene.radius = 450
	if font, err := ttf.OpenFont("DejaVuSans.ttf", 18); err != nil {
		panic(err)
	} else {
		scene.sans18 = font
	}
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	window.SetTitle("Scene")
}

// Draw draws a single frame of the scene
func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	var x1, y1, x2, y2 float64
	cx, cy := float64(ScreenWidth/2), float64(ScreenHeight/2)
	for n := 0; n < scene.points; n++ {
		n1 := float64(n)
		n2 := scene.factor * n1
		theta := 2 * math.Pi * n1 / float64(scene.points)
		phi := 2 * math.Pi * n2 / float64(scene.points)
		x1 = scene.radius * math.Cos(theta)
		y1 = scene.radius * math.Sin(theta)
		x2 = scene.radius * math.Cos(phi)
		y2 = scene.radius * math.Sin(phi)
		gfx.LineColor(renderer,
			int32(cx-x1), int32(cy-y1),
			int32(cx-x2), int32(cy-y2), Green)
	}

	gfx.CircleColor(renderer, 200, 100, 25, Green)
	gfx.RectangleColor(renderer, 100, 100, 200, 50, Red)

	factor := fmt.Sprintf("factor: %6.3f", scene.factor)
	gfx.StringColor(renderer, 20, 20, factor, Yellow)
}
