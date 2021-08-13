package main

import "encoding/base64"

func main() {
	uustring := "`!\"#$%&'()*+,-./0123456789:;<|>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_"
	uustring2 := "`!\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_"
	uuencode := base64.NewEncoding(uustring)
	uuencode2 := base64.NewEncoding(uustring2)

	println("The following two strings should be equal")
	ret, err := uuencode.DecodeString("4$114$||")
	println(string(ret))
	if err != nil {
		panic(err)
	}
	ret, err = uuencode2.DecodeString("4$114$==")
	println(string(ret))
	if err != nil {
		panic(err)
	}

	println("The following two strings should be equal")
	ret, err = uuencode.DecodeString("4$114$|1")
	println(string(ret))
	if err != nil {
		panic(err)
	}
	ret, err = uuencode2.DecodeString("4$114$=1")
	println(string(ret))
	if err != nil {
		panic(err)
	}
}
