package piscine

func Abort(a, b, c, d, e int) int {
	arg := []int{a, b, c, d, e}     // create a slice of int
	var minArr []int                // create a slice of int
	for i := 0; i < len(arg); i++ { // loop through the slice
		for j := 1; j < len(arg)-i; j++ { // loop through the slice
			if arg[j] < arg[j-1] { // compare the values
				arg[j], arg[j-1] = arg[j-1], arg[j] // swap the values
			}
		}
		minArr = append(minArr, arg[i]) // append the minimum value to the slice
	}
	return minArr[2] // return the third minimum value
}
