package piscine

func TrimAtoi(s string) int {
	res1 := 0
	res2 := 1
	for _, item := range s {
		if item >= '0' && item <= '9' {
			res1 = (int(item) - '0') + res1*10
		} else if item == '-' && res1 == 0 {
			res2 = -1
		}
	}
	return res1 * res2
}
