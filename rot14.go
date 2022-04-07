package piscine

import "github.com/01-edu/z01"

func Rot14(s string) string {
	a := []rune(s)
	for _, i := range a {
		if rune(a[i+14]) > 'Z' && rune(a[i+14]) < 'a' || rune(a[i+14]) > 'z' {
			z01.PrintRune(rune(a[i-12]))
		} else {
			z01.PrintRune(rune(a[i+14]))
		}
	}
	return string(a)
}
