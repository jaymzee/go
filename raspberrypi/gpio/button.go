package gpio

import (
	"fmt"
	"io"
	"os"
)

type Button struct {
	pin  int
	file *os.File
}

func NewButton(pin int) *Button {
	filename := fmt.Sprintf("/sys/class/gpio/gpio%d/value", pin)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	return &Button{pin: pin, file: file}
}

func (btn *Button) Pressed() bool {
	buffer := make([]byte, 16)

	btn.file.Seek(0, io.SeekStart)
	count, err := btn.file.Read(buffer)
	if err != nil {
		panic(err)
	}

	return !(count > 0 && buffer[0] == '1')
}
