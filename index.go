package piscine

func Index(s string, toFind string) int {
	length := len(s)
	index := len(toFind)

	t := 0
	for i := 0; i < length; i++ {
		j := 0
		t = i
		for j < index {
			if t < length && s[t] == toFind[j] {
				j++
				t++
			} else {
				break
			}
		}
		if j == index {
			return i
		}
	}
	return -1
}
