package piscine

func IsUpper(s string) bool {
	str := []rune(s)
	for i := range str {
		if (str[i] < 'A') && (str[i] > 'Z') {
			return false
		}
	}
	return true
}
