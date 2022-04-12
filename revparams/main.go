package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	x := 0
	for i := range arg {
		x = i
	}
	for e := x; e > 0; e-- {
		for _, j := range arg[e] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
