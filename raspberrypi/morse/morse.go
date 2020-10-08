package main

import (
	"github.com/jaymzee/go/raspberrypi/gpio"
	"log"
	"strings"
	"time"
)

var morse = map[rune]string{
	'A': ".-",
	'B': "-...",
	'C': "-.-.",
	'D': "-..",
	'E': ".",
	'F': "..-.",
	'G': "--.",
	'H': "....",
	'I': "..",
	'J': ".---",
	'K': "-.-",
	'L': ".-..",
	'M': "--",
	'N': "-.",
	'O': "---",
	'P': ".--.",
	'Q': "--.-",
	'R': ".-.",
	'S': "...",
	'T': "-",
	'U': "..-",
	'V': "...-",
	'W': ".--",
	'X': "-..-",
	'Y': "-.--",
	'Z': "--..",
	'1': ".----",
	'2': "..---",
	'3': "...--",
	'4': "....-",
	'5': ".....",
	'6': "-....",
	'7': "--...",
	'8': "---..",
	'9': "----.",
	'0': "-----",
}

func dit(led *gpio.LED) {
	led.Out(0) // turn on LED
	time.Sleep(100 * time.Millisecond)
	led.Out(1) // turn off LED
	time.Sleep(100 * time.Millisecond)
}

func dah(led *gpio.LED) {
	led.Out(0)
	time.Sleep(300 * time.Millisecond)
	led.Out(1)
	time.Sleep(100 * time.Millisecond)
}

func sendChar(led *gpio.LED, ch rune) {
	if code, found := morse[ch]; found {
		log.Printf("%q %c\n", code, ch)
		for _, sym := range code {
			switch sym {
			case '.':
				dit(led)
			case '-':
				dah(led)
			default:
				panic("invalid character in morse code table")
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
	// ignore characters not in morse map
}

func sendString(led *gpio.LED, msg string) {
	for _, ch := range strings.ToUpper(msg) {
		switch ch {
		case ' ':
			log.Println("SPC")
			time.Sleep(400 * time.Millisecond)
		case '.':
			log.Println("STOP")
			time.Sleep(1 * time.Second)
		default:
			sendChar(led, ch)
		}
	}
}
