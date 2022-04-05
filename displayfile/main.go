package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	arguments := os.Args
	if len(arguments) == 2 {
		fileName := arguments[1]
		content, err := ioutil.ReadFile(fileName)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Print(string(content))
		fmt.Println()
	} else if len(arguments) > 2 {
		fmt.Println("Too many arguments")
	} else {
		fmt.Println("File name missing")
	}
}
