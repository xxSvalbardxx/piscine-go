package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	x := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		x[i] = i + min
	}
	return x
}
