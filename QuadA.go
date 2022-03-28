package main

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	for h := 1; h <= y; h++ {
		for w := 1; w <= x; w++ {
			if h == 1 && w == 1 || h == y && w == x || h == 1 && w == x || h == y && w == 1 {
				z01.PrintRune('o')
			} else if w == 1 || w == x {
				z01.PrintRune('|')
			} else if h == 1 && w != 1 && w != x || h == y && w != 1 && w != x {
				z01.PrintRune('-')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	QuadA(-1, -2)
}
