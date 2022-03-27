package piscine

import "github.com/01-edu/z01"

func Printstr(s string) {
	s = "Hello World!"
	for s := 0; s < 12; s++ {
		z01.PrintRune(Rune(s[0]))
		z01.PrintRune('\n')
	}
}
