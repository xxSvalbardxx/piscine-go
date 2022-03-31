package piscine

func TrimAtoi(s string) int {
	res1 := 0
	res2 := 1
	for _, item := range s { // "_" remplace index et "item" est la copie du "_" | range s pour aller chercher tout les chiffres qu'on demande
		if item >= '0' && item <= '9' { // Si l'item est compris entre 0 et 9
			res1 = (int(item) - '0') + res1*10 // Prendre l'item puis le mutliplié par 10 pour garder ceux au dessus de 0 et 9
		} else if item == '-' && res1 == 0 { // Sinon ceux qui commence par '-' ou 0
			res2 = -1 // -1 pour mettre le moins devant et ne pas prendre les 0
		}
	}
	return res1 * res2
}
