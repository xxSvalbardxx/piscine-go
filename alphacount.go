package piscine

func AlphaCount(s string) int {
	x := 0
	str := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if (str[i] >= 'A') && (str[i] <= 'Z') || (str[i] > 'a') && (str[i] < 'z') {
			x = x + 1
		}
	}
	return x
}
