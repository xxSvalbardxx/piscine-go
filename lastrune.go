package piscine

func LastRune(s string) rune {
	/*	last := []rune(s)
		length := 0
		for i := range s {
			length = i + 1
		}
		return last[length-1]
	}*/
	a := []rune(s)
	return a[len(s)-1]
}
