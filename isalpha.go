package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if (str[i] < '0') || (str[i] > '9') && (str[i] < 'A') || (str[i] > 'Z') && (str[i] < 'a') || (str[i] > 'z') {
			return false
		}
	return true
}