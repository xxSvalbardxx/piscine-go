package piscine

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else {
		return Fibonacci(index-1) + (index - 2)
	}
}
