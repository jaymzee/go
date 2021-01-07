package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"math"
	"math/rand"
)

const (
	// ScreenWidth of window
	ScreenWidth = 924
	// ScreenHeight of window
	ScreenHeight = 800
	// FPS is frames per second
	FPS = 10
	PHI = (1 + 2.23606797749979) / 2
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
	window.SetTitle("chaos")
}

func (scene *Scene) tick(window *sdl.Window, renderer *sdl.Renderer) {
	scene.counter++
}

type Point2D struct {
	X int32
	Y int32
}

type Shape struct {
	VertexList []Point2D
	Factor     float64
}

var triangle = Shape{
	VertexList: []Point2D{
		Point2D{ScreenWidth / 2, 0},
		Point2D{0, ScreenHeight},
		Point2D{ScreenWidth, ScreenHeight},
	},
	Factor: 1 / 2.0,
}

var vicsek = Shape{
	VertexList: []Point2D{
		Point2D{0, 0},
		Point2D{ScreenWidth, 0},
		Point2D{ScreenWidth, ScreenHeight},
		Point2D{0, ScreenHeight},
		Point2D{ScreenWidth / 2, ScreenHeight / 2},
	},
	Factor: 2 / 3.0,
}

var carpet = Shape{
	VertexList: []Point2D{
		Point2D{0, 0},
		Point2D{ScreenWidth, 0},
		Point2D{ScreenWidth, ScreenHeight},
		Point2D{0, ScreenHeight},
		Point2D{ScreenWidth / 2, 0},
		Point2D{ScreenWidth, ScreenHeight / 2},
		Point2D{ScreenWidth / 2, ScreenHeight},
		Point2D{0, ScreenHeight / 2},
	},
	Factor: 2 / 3.0,
}

func CreatePentagon() Shape {
	var vlist [5]Point2D
	w := float64(ScreenWidth)
	h := float64(ScreenHeight)
	for n := 0; n < 5; n++ {
		x := w/2.0 + h/2.0*math.Sin(float64(n)*2.0*math.Pi/5.0)
		y := h/20.0 + h/2.0*(1.0-math.Cos(float64(n)*2.0*math.Pi/5.0))
		vlist[n] = Point2D{int32(x), int32(y)}
	}
	return Shape{VertexList: vlist[:], Factor: 1 / PHI}
}

func sierpinski(p Point2D, vlist []Point2D, factor float64) Point2D {
	// select a vertex at random
	n := int(float64(len(vlist)) * rand.Float64())
	v := vlist[n]
	// move a factor of the distance to that vertex
	dx := float64(v.X-p.X) * factor
	dy := float64(v.Y-p.Y) * factor
	x := p.X + int32(dx)
	y := p.Y + int32(dy)
	return Point2D{x, y}
}

func (scene *Scene) draw(window *sdl.Window, renderer *sdl.Renderer) {
	renderer.SetDrawColor(0, 0, 0, 255)
	renderer.Clear()
	renderer.SetDrawColor(0, 255, 0, 255)
	shape := &carpet
	p := Point2D{ScreenWidth / 2, ScreenHeight / 2}
	for i := 0; i < 10000; i++ {
		if i > 5 {
			// we throw away the first few points
			renderer.DrawPoint(p.X, p.Y)
		}
		p = sierpinski(p, shape.VertexList, shape.Factor)
	}
}
