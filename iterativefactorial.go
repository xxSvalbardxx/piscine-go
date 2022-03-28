package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else if nb == nb*(nb-1) {
		return nb
	}
}
