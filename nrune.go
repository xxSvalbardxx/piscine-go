package piscine

func NRune(s string, n int) rune {
	i := []rune(s)
	length := 0
	for index := range i {
		length = index
	}

	if n-1 >= 0 && n-1 <= length {
		return i[n-1]
	}
	return 0

