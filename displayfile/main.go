package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arg := os.Args
	if len(arg) == 2 { // L'arg[0]= nom du dossier "displayfile", l'arg[1]= nom du fichier
		fileName := arg[1]
		content, err := ioutil.ReadFile(fileName)
		if err != nil { // obligatoire pour que ioutil fonctionne
			fmt.Println("Error")
		}
		fmt.Println(string(content))
	} else if len(arg) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
}
