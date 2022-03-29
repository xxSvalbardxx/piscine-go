package piscine

func NRune(s string, n int) rune {
	r := []rune(s)
	len := 0
	for index := range r {
		len = index
	}

	if n-1 >= 0 && n-1 <= len {
		return r[n-1]
	}
	return 0
}
