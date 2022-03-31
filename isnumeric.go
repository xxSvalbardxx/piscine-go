package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] < '0' || str[i] > '9' {
			return false
		}
	}
	return true
}
