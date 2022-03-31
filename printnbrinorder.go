package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	my_arr := [10]int{}

	if n < 10 && n >= 0 {
		z01.PrintRune(rune(n + 48))
	} else {

		for {
			if n == 0 {
				break
			}
			my_arr[n%10]++
			n /= 10
		}

		for index, num := range my_arr {
			if num != 0 {
				for i := 0; i < num; i++ {
					z01.PrintRune(rune(index + 48))
				}
			}
		}			//z01
	}
}
