package main

import (
	"fmt"
	"github.com/jaymzee/go/sdl/ssd"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"time"
)

const (
	// ScreenWidth of window
	ScreenWidth = 400
	// ScreenHeight of window
	ScreenHeight = 400
)

// Colors
var (
	Red    = sdl.Color{R: 255, G: 0, B: 0, A: 255}
	Green  = sdl.Color{R: 0, G: 255, B: 0, A: 255}
	Yellow = sdl.Color{R: 255, G: 255, B: 0, A: 255}
)

// Scene contains the state for the scene
type Scene struct {
	ch1    chan int
	ch2    chan int
	sans18 *ttf.Font
	ssd1   ssd.Display
	ssd2   ssd.Display
}

// Loop is the event loop for the scene
func (scene *Scene) Loop(window *sdl.Window, renderer *sdl.Renderer) {
	scene.init(window, renderer)
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
	}
}

// NewScene initializes the scene
func (scene *Scene) init(window *sdl.Window, renderer *sdl.Renderer) {
	if font, err := ttf.OpenFont("DejaVuSans.ttf", 18); err != nil {
		panic(err)
	} else {
		scene.sans18 = font
	}
	if disp, err := ssd.OpenSSD("ssd10d75.json"); err != nil {
		panic(err)
	} else {
		scene.ssd1 = *disp
		scene.ssd2 = *disp
		scene.ssd1.OnColor = Red
		scene.ssd2.OnColor = Green
	}

	window.SetTitle("chan")

	scene.ch1 = make(chan int)
	go ticker(scene.ch1, 1000*time.Millisecond)
	scene.ch2 = make(chan int)
	go ticker(scene.ch2, 16666*time.Microsecond)
}

// Draw draws a single frame of the scene
func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	var i int32
	select {
	case number := <-scene.ch1:
		for i = 0; i < 4; i++ {
			digit := number >> (4 * i) & 0xF
			b := scene.ssd1.Encode(digit, false)
			scene.ssd1.Draw(renderer, 250-50*i, 100, b)
		}
	case number := <-scene.ch2:
		for i = 0; i < 4; i++ {
			digit := number >> (4 * i) & 0xF
			b := scene.ssd2.Encode(digit, false)
			scene.ssd2.Draw(renderer, 250-50*i, 200, b)
		}
	}
}

func ticker(ch chan int, interval time.Duration) {
	for i := 0; ; i++ {
		ch <- i
		time.Sleep(interval)
	}
}
