package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	s = "Hello World!"
	z01.PrintRune(%p, s)
}
