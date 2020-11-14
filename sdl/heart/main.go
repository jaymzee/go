package main

import (
	"flag"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
)

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
	window, err := sdl.CreateWindow("Initializing SDL...",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		ScreenWidth, ScreenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, rendererFlags)
	if err != nil {
		panic(err)
	}
	defer renderer.Destroy()

	// init scene
	var scene Scene = new(Heart)
	scene.Init(window, renderer)
	// scene render loop
	for running := true; running; {
		scene.Draw(window, renderer)
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
