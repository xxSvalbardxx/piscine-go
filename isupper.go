package piscine

func IsUpper(s string) bool {
	x := 0
	str := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if (str[i] >= 'A') && (str[i] <= 'Z') || (str[i] >= 'a') && (str[i] <= 'z') {
			x += 1
			return true
		}
	}
	return false
}
