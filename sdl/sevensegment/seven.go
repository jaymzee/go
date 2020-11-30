package sevensegment

import (
	"encoding/json"
	"errors"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

// Display is the type for configuring seven-segment display settings
type Display struct {
	Code        [17]uint8
	Border      sdl.Rect
	X           [7][6]int16
	Y           [7][6]int16
	P           sdl.Point
	PSize       int32
	FillColor   sdl.Color
	BorderColor sdl.Color
	LedOnColor  sdl.Color
	LedOffColor sdl.Color
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
	vp := d.Border
	vp.X = x
	vp.Y = y

	// fill shape with background color (erase it)
	fill := d.FillColor
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
		led := d.LedOffColor
		if (b>>i)&1 == 1 {
			led = d.LedOnColor
		}
		if i < 7 {
			if !gfx.FilledPolygonColor(rend, d.X[i][:], d.Y[i][:], led) {
				return errors.New("draw seven-segment led failed")
			}
		} else {
			if !gfx.FilledCircleColor(rend, d.P.X, d.P.Y, d.PSize, led) {
				return errors.New("draw seven-segment led failed")
			}
		}
	}

	return nil
}

// Encode encodes a digit for a seven-segment display
// the digit may be hexadecimal, decimal, or octal
func (d *Display) Encode(digit int, point bool) uint8 {
	code := d.Code[16] // indicate an error
	if digit >= 0 && digit < 16 {
		code = d.Code[digit]
	}
	if point {
		code |= 0x80
	}
	return code
}
