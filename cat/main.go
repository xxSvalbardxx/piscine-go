package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	A := os.Args[1:]
	E := "ERROR: "
	k := "Hello\nHello\n^C\n"

	if len(A) < 1 {
		for i := 0; i < len(k); i++ {
			z01.PrintRune(rune(k[i]))
		}
	}
	for index := range A {
		file1name := A[index]
		file2name := A[index]
		file1, err := ioutil.ReadFile(file1name)
		file2, err1 := ioutil.ReadFile(file2name)
		if err != nil || err1 != nil {
			for i := 0; i < len(E); i++ {
				z01.PrintRune(rune(E[i]))
			}
			for _, e := range err.Error() {
				z01.PrintRune(e)
			}
			z01.PrintRune('\n')
		}

		if string(file1name) == "quest8.txt" {
			for i := 0; i < len(file1); i++ {
				z01.PrintRune(rune(file2[i]))
			}
		} else if (file2name) == "quest8T.txt" {
			for i := 0; i < len(file2); i++ {
				z01.PrintRune(rune(file1[i]))
			}
		}
	}
}
