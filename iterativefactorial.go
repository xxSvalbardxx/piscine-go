package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb = 1 || nb = 0 {
		return 1
	} else {
		for i:=1; i<=nb; i++ {
			nb = nb*(nb-1)
			return i
	}
}