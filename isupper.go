package piscine

func IsUpper(s string) bool {
	x := 0
	str := []rune(s)
	for i := range str {
		if (str[i] < 'A') && (str[i] > 'Z') {
			x += 1
			if x > 0 {
				return false
			}
			continue
		}
	}
	return true
}
