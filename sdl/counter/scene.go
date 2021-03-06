package main

import (
	"fmt"
	"github.com/jaymzee/go/sdl/ssd"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

const (
	// ScreenWidth of window
	ScreenWidth = 400
	// ScreenHeight of window
	ScreenHeight = 300
	// FPS is frames per second
	FPS = 4
)

// Colors
var (
	Red    = sdl.Color{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}
	Green  = sdl.Color{R: 0x00, G: 0xFF, B: 0x00, A: 0xFF}
	Yellow = sdl.Color{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF}
)

// Scene contains the state for the scene
type Scene struct {
	counter int
	sans18  *ttf.Font
	ssd     *ssd.Display
	colon   *ssd.Colon
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
				//fmt.Printf("%T %#v\n", event, event)
			}
		}

		// draw a single frame of the scene
		scene.draw(window, renderer)
		renderer.Present()
		gfx.FramerateDelay(fpsmgr)

		// update scene
		scene.tick(window, renderer)
	}
}

// NewScene initializes the scene
func (scene *Scene) init(window *sdl.Window, renderer *sdl.Renderer) {
	scene.counter = 0
	if font, err := ttf.OpenFont("DejaVuSans.ttf", 18); err != nil {
		panic(err)
	} else {
		scene.sans18 = font
	}
	if disp, err := ssd.OpenSSD("ssd10d75.json"); err != nil {
		panic(err)
	} else {
		scene.ssd = disp
	}
	window.SetTitle("seven-segment display")
}

// Draw draws a single frame of the scene
func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()

	code := scene.ssd.Encode(scene.counter&0xF, false)
	if err := scene.ssd.Draw(renderer, 100, 100, code); err != nil {
		panic(err)
	}

	counterText := fmt.Sprintf("counter: %#x", scene.counter)
	DrawText(renderer, 20, 20, counterText, scene.sans18, Yellow)
	codeText := fmt.Sprintf("encoded: %08b", code)
	DrawText(renderer, 20, 40, codeText, scene.sans18, Yellow)
}

func (scene *Scene) tick(window *sdl.Window, renderer *sdl.Renderer) {
	scene.counter++
}
