package piscine

func IsPrime(nb int) bool {
	rslt := true
	if nb == 0 || nb == 1 {
		return false
	} else {
		for i := 2 ; i <= nb / 2 ;i++ {
			if nb%i == 0 {
				return false
				rslt = false
				break
			}
		}
			if rslt == true
				return true
	}
}

