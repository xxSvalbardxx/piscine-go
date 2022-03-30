package piscine

func IsAlpha(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if (str[i] < 'A') || (str[i] > 'Z') {
			return false
		} else if (str[i] < 'a') || (str[i] > 'z') {
			return false
		} else if (str[i] < '0') || (str[i] > '9') {
			return false
		}
	return true
}