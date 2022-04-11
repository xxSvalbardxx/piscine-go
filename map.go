package piscine

func Map(f func(int) bool, a []int) []bool {
	for _, nbr := range a {
		f(nbr)
	}
}
func IsPrime(nb int) bool {}
