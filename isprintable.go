package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for i := 0; i < len(str); i++ {
		if str[i] < 32 {
			return false
		}
	}
	return true
}
