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
	for i := x; i > 0; i-- {
		for _, j := range arg[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
