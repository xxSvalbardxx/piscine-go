package piscine

func Rot14(s string) string {
	r := []rune(s)
	var result string
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i] <= 'z' {
			if r[i] >= 'm' {
				r[i] = r[i] - 12
			} else {
				r[i] = r[i] + 14
			}
		} else if r[i] >= 'A' && r[i] <= 'Z' {
			if r[i] >= 'M' {
				r[i] = r[i] - 12
			} else {
				r[i] = r[i] + 14
			}
		}
		result += string(r[i])
	}
	return result
}
