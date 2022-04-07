package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) { // Découpage en Runes de la phrase de fin, puisque seul PrintRune est authorisé.
	for _, r := range s { // Imprime les runes une a une.
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 { // Calcule le reste de: Nbr d'arg / 2
		return true
	} else {
		return false
	}
}

func main() {
	arg := os.Args[1:]    // Prends en compte les Args. A partir du premier ( os.Arg[0] = Nom du Fichier ).
	if isEven(len(arg)) { // si la longeur (le nbr) des args correspond a la condit° isEven
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
