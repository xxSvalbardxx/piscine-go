package piscine

func Sqrt(nb int) int {
	if nb > 0 {
		resultat, sqrt := 1, 0
		for i := 1; resultat <= nb; i++ {
			resultat = i * i
			sqrt++
			if resultat == nb {
				return sqrt
			}
		}
		return 0
	} else {
		return 0
	}
}
