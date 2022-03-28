package piscine

func Sqrt(nb int) int {
	result := 0
	if nb < 0 {
		return 0
	} else {
		result = nb / Sqrt(nb)
	}
	return result
}
