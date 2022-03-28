package piscine

func IterativePower(nb int, power int) int {
	if nb <= 0 {
		return 0
	} else if nb >= 1 && nb <= 2147483647 {
		return nb * nb
	} else {
		return 1
	}
}
