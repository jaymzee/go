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

// Heart contains the state for our scene
type Heart struct {
	Points int
	Factor float64
	Radius float64
	Sans18 *ttf.Font
}

// Init initializes the scene
func (heart *Heart) Init(window *sdl.Window, renderer *sdl.Renderer) {
	var err error
	heart.Factor = 1.0
	heart.Points = 500
	heart.Radius = 450
	heart.Sans18, err = ttf.OpenFont("DejaVuSans.ttf", 18)
	if err != nil {
		panic(err)
	}
	renderer.SetDrawBlendMode(sdl.BLENDMODE_BLEND)
	window.SetTitle("Heart")
}

// Draw draws a single frame of the scene
func (heart *Heart) Draw(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 64)
	var x1, y1, x2, y2 float64
	cx, cy := float64(ScreenWidth/2), float64(ScreenHeight/2)
	for n := 0; n < heart.Points; n++ {
		n1 := float64(n)
		n2 := heart.Factor * n1
		theta := 2 * math.Pi * n1 / float64(heart.Points)
		phi := 2 * math.Pi * n2 / float64(heart.Points)
		x1 = heart.Radius * math.Cos(theta)
		y1 = heart.Radius * math.Sin(theta)
		x2 = heart.Radius * math.Cos(phi)
		y2 = heart.Radius * math.Sin(phi)
		renderer.DrawLine(
			int32(cx-x1), int32(cy-y1),
			int32(cx-x2), int32(cy-y2),
		)
	}

	factor := fmt.Sprintf("factor: %6.3f", heart.Factor)
	DrawText(renderer, factor, 200, 200, heart.Sans18, Yellow)
	heart.Factor += 0.001
}
