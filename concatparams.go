package piscine

func ConcatParams(args []string) string {
	res := "" // Résultat est égal à un string vide (sale)

	for _, val := range args { // Pour la valeur et non l'index dans ce que contient args
		res += "\n" + val // resultat est égal à retour à la ligne + valeur
	}
	return res[1:] // On retourne le résultat, de l'index 1 (: permet d'appeler un string)
}
