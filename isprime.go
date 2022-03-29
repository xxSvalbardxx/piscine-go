package piscine

func IsPrime(nb int) bool {
	for (let i = 2; i * i <= nb; i++) {
		if (nb % i == 0) {
	  	return false 
	  	}
	}
	return nb > 1;
}
