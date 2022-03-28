package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 {
		return 0
	} else if power == 0 {
		return 1
	} else {
		result = nb * (IterativePower(nb, power-1))
	}
	return result
}
