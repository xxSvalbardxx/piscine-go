package piscine

func Rot14(s string) string {
	runes := []rune(s)
	result := make([]rune, 0)

	for _, l := range runes {
		if l >= 'a' && l <= 'z' {
			if l+14 > 'z' {
				result = append(result, 'a'+l-'a'+14-26)
			} else {
				result = append(result, 'a'+l-'a'+14)
			}
		} else if l >= 'A' && l <= 'Z' {
			if l+14 > 'Z' {
				result = append(result, 'A'+l-'A'+14-26)
			} else {
				result = append(result, 'A'+l-'A'+14)
			}
		} else {
			result = append(result, l)
		}
	}

	return string(result)
}
