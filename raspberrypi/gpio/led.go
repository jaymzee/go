package gpio

import (
	"os"
	"fmt"
)

type LED struct {
	pin int
	file *os.File
}

func NewLED(pin int) *LED {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)
	file, err := os.OpenFile(filename, os.O_WRONLY, 0)
	if err != nil {
		panic(err)
	}
	return &LED{pin: pin, file: file}
}

func (led *LED) Set(bit byte) {
	led.file.Write([]byte{'0' + bit})
}
