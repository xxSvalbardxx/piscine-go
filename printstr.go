package piscine

import "github.com/01-edu/z01"

func Printstr(s string) {
	s = "Hello World!"
	for s := 0; s <= 12; s++ {
		z01.PrintRune(rune(s))
		z01.PrintRune('\n')
	}
}
