package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a)) // create a new slice of bool, with the same length as the original slice

	for i, e := range a { // iterate over the original slice, and apply the function to each element
		tab[i] = f(e) // assign the result of the function to the new slice, at the same index
	}
	return tab
}
