package piscine

func SplitWhiteSpaces(s string) []string {
	var printstr []string // creer la nouvelle variable
	j := 0
	for i := 0; i < len(s)-1; i++ {
		if rune(s[i]) == ' ' { // a chaque fois que la boucle tombe sur un espace
			if i != j { //
				printstr = append(printstr, s[j:i])
			}
			j = i + 1 // permet d'eviter la repetition du mot precedent
		} else if i == len(s)-4 { //-1 / - 4 permet d'avoir toute la Phrase
			printstr = append(printstr, s[j:])
		}
	}
	return printstr
}
