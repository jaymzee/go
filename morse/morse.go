package morse

import (
	"log"
	"strings"
	"time"
)

var Code = map[rune]string{
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

type Output interface {
	On()
	Off()
}

func Dit(o Output) {
	o.On()
	time.Sleep(100 * time.Millisecond)
	o.Off()
	time.Sleep(100 * time.Millisecond)
}

func Dah(o Output) {
	o.On()
	time.Sleep(300 * time.Millisecond)
	o.Off()
	time.Sleep(100 * time.Millisecond)
}

func SendChar(o Output, ch rune) {
	if code, found := Code[ch]; found {
		log.Printf("%c %q\n", ch, code)
		for _, sym := range code {
			switch sym {
			case '.':
				Dit(o)
			case '-':
				Dah(o)
			default:
				panic("invalid character in morse code table")
			}
		}
		time.Sleep(200 * time.Millisecond)
	}
	// ignore characters not in morse map
}

func Send(o Output, msg string) {
	for _, ch := range strings.ToUpper(msg) {
		switch ch {
		case ' ':
			log.Println("SPC")
			time.Sleep(400 * time.Millisecond)
		case '.':
			log.Println("STOP")
			time.Sleep(1 * time.Second)
		default:
			SendChar(o, ch)
		}
	}
}
