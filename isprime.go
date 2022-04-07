package piscine

func IsPrime(nb int) bool {
	if nb == 0 || nb == 1 || nb < 0 {
		return false
	} else {
		for i := 2; i <= nb/2; i++ {
			if nb%i == 0 { // si le nombre est divisible et que le resultat est un nbr a virgule, alors le nbr n'est pas premier
				return false
			}
		}
	}
	return true
}
