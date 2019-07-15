package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// blank line by default
	var str string

	// with special characters
	var hello string = "Hello\n\t"

	// without special characters
	var world string = `There\n\t`

	fmt.Println("str", str)
	fmt.Println("hello", hello)
	fmt.Println("world", world)

	// UTF-8
	var helloWorld = "Hello There!"
	hi := "你好，世界"

	fmt.Println("helloWorld", helloWorld)
	fmt.Println("hi", hi)

	// single quotes for bytes (uint8)
	var rawBinary byte = '\x27'

	// rune (uint32) for UTF-8
	var someChinese rune = '茶'

	fmt.Println(rawBinary, someChinese)

	helloWorld = "Hello There"
	// string concatenation
	andGoodMorning := helloWorld + " and Others!"

	fmt.Println(helloWorld, andGoodMorning)

	// strings are immutable
	// cannot assign to helloWorld[0]
	// helloWorld[0] = 72

	// getting the string length
	byteLen := len(helloWorld)                    // 19 byte
	symbols := utf8.RuneCountInString(helloWorld) // 10 rune

	fmt.Println(byteLen, symbols)

	// getting the substring
	hello = helloWorld[:12] 
	H := helloWorld[0]      
	fmt.Println(H)

	// convert to slice byte and back
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)

	fmt.Println(byteString, helloWorld)
}
