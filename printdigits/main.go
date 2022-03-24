package main

import "github.com/01-edu/z01"

func main() {
	for i := '1'; i <= '9'; i++ {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
