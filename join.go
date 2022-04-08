package piscine

func Join(strs []string, sep string) string {
	x := ""
	for i, val := range strs {
		if i != len(strs)-1 {
			x += val + sep
		} else {
			x += val
		}
	}
	return x
}
