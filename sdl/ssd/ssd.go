package ssd

import (
	"encoding/json"
	"errors"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

// Display holds the configuration for the seven-segment display
type Display struct {
	Encoder     [17]uint8
	Size        sdl.Rect
	X           [7][6]int16
	Y           [7][6]int16
	P           sdl.Point
	PR          int32
	FaceColor   sdl.Color
	BorderColor sdl.Color
	OnColor     sdl.Color
	OffColor    sdl.Color
}

// Open returns a Display configured from file
func Open(name string) (*Display, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	bytes := make([]byte, info.Size())
	_, err = file.Read(bytes)
	if err != nil {
		return nil, err
	}
	disp := new(Display)
	err = json.Unmarshal(bytes, disp)
	if err != nil {
		return nil, err
	}
	return disp, nil
}

// Write saves display settings to file
func (d *Display) Write(name string) error {
	b, err := json.Marshal(d)
	if err != nil {
		return err
	}
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(b)
	return nil
}

// Draw draws a seven-segment display on the screen
//  renderer is the renderer to draw on
//  x is the x coordinate of the upper left corner
//  y is the y coordinate of the upper left corner
//  b drives the A,B,C,D,E,F,G leds (bit-0..bit-6 -> A..G, bit-7 -> dot)
//  c is the color of the on state of the leds
func (d *Display) Draw(rend *sdl.Renderer, x, y int32, b uint8) error {
	// calculate new viewport
	vp := sdl.Rect{x, y, d.Size.X, d.Size.Y}

	// fill shape with background color (erase it)
	fill := d.FaceColor
	if err := rend.SetDrawColor(fill.R, fill.G, fill.B, fill.A); err != nil {
		return err
	}
	if err := rend.FillRect(&vp); err != nil {
		return err
	}

	// save the current viewport and set to the new one
	vpsave := rend.GetViewport()
	defer rend.SetViewport(&vpsave)
	if err := rend.SetViewport(&vp); err != nil {
		return err
	}

	// draw border
	if !gfx.RectangleColor(rend, 0, 0, vp.W, vp.H, d.BorderColor) {
		return errors.New("draw seven-segment border failed")
	}

	// draw leds
	for i := 0; i < 8; i++ {
		led := d.OffColor
		if (b>>i)&1 == 1 {
			led = d.OnColor
		}
		if i < 7 {
			if !gfx.FilledPolygonColor(rend, d.X[i][:], d.Y[i][:], led) {
				return errors.New("draw seven-segment led failed")
			}
		} else {
			if !gfx.FilledCircleColor(rend, d.P.X, d.P.Y, d.PR, led) {
				return errors.New("draw seven-segment led failed")
			}
		}
	}

	return nil
}

// Encode encodes a digit for a seven-segment display
// the digit may be hexadecimal, decimal, or octal
func (d *Display) Encode(digit int, point bool) uint8 {
	code := d.Encoder[16] // default to indicating an error
	if digit >= 0 && digit < 16 {
		code = d.Encoder[digit]
	}
	if point {
		code |= 0x80
	}
	return code
}
