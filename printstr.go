package piscine

import "github.com/01-edu/z01"

func Printstr(s string) {
	s = "Hello World!"
	z01.PrintRune(s)
	z01.PrintRune('\n')
}
