package main

import "fmt"

func SplitWhiteSpaces(s string) []string {
	var printstr []string
	j := 0
	for i := 0; i < len(s)-1; i++ {
		if rune(s[i]) == ' ' {
			if i != j {
				printstr = append(printstr, s[j:i])
			}
			j = j
		} else if i == len(s)-4 { //-1 / - 4 permet d'avoir toute la Phrase
			printstr = append(printstr, s[j:])
		}
	}
	return printstr
}

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}
