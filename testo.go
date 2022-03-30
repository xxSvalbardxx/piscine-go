package main

import "github.com/01-edu/z01."

func main() {
	str := "Hello"
	length := 0
	for i := range str {
		length = i + 1
		z01.PrintRune(rune(length))
	}
	return length
}
