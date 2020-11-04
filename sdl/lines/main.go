package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
)

const fps = 60

func draw(r *sdl.Renderer, f *ttf.Font, t float64) {
	yellow := sdl.Color{255, 255, 0, 255}

	r.SetDrawColor(0, 0, 0, 255)
	r.Clear()

	hellosurf, err := f.RenderUTF8Solid("hello", yellow)
	if err != nil {
		panic(err)
	}
	hello, err := r.CreateTextureFromSurface(hellosurf)
	if err != nil {
		panic(err)
	}
	r.Copy(hello, nil, &sdl.Rect{20, 20, 40, 20})
	x := int32(200 * math.Cos(-2.0*math.Pi*t))
	y := int32(200 * math.Sin(-2.0*math.Pi*t))
	r.SetDrawColor(0, 255, 0, 255)
	r.DrawLine(400, 300, 400+x, 300-y)
	r.Present()
}

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	if err := ttf.Init(); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer(800, 600, 0)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()
	window.SetTitle("Hello World!")

	font, err := ttf.OpenFont("FreeSans.ttf", 24)
	if err != nil {
		panic(err)
	}

	t := 0.0
	for running := true; running; {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			}
		}
		draw(renderer, font, t)
		sdl.Delay(1000 / fps)
		t += 1.0 / fps
	}
}
