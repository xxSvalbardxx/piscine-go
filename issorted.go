package piscine

func IsSorted(f func(a, b int) int, tab []int) bool {
	if len(tab) > 1 { // if the slice is not empty
		if f(tab[0], tab[1]) >= 0 { // if the first element is greater than or equal to the second element, return false
			for i := 0; i < len(tab)-1; i++ { // iterate over the slice
				if f(tab[i], tab[i+1]) < 0 { // if the function returns a negative value, the slice is not sorted
					return false
				}
			}
		}
		if f(tab[0], tab[1]) <= 0 {
			for i := 0; i < len(tab)-1; i++ {
				if f(tab[i], tab[i+1]) > 0 {
					return false
				}
			}
		}
	}
	return true
}
