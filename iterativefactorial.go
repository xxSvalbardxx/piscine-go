package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		rslt := 1
		for i := 1; i <= nb; i++ {
			rslt *= i
			if nb > 2147483647 {
				return 0
			}
		}
	}
}
