package piscine

func Unmatch(a []int) int {
	var un int

	for _, x := range a {
		un = 0
		for _, y := range a {
			if y == x {
				un++
			}
		}
		if un%2 != 0 {
			return x
		}
	}
	return -1
}
