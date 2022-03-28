package piscine

func RecursiveFactorial(nb int) int {
	if nb == 0 || nb == 1 {
		return 1
	} else if nb < 0 {
		return 0
	} else if nb > 1 && nb > 2147483647 {
		return nb * RecursiveFactorial(nb-1)
	}
}
