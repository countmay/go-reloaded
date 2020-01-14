package main

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	a := os.Args[1:]
	l := 0

	for range a {
		l++
	}
	if l == 0 {

		io.Copy(os.Stdin, os.Stdout)

	}
	for i := 0; i < l; i++ {

		content, err := ioutil.ReadFile(a[i])
		if err != nil {
			er := "open " + a[i] + " : no such file or directory"
			for _, g := range er {
				z01.PrintRune(rune(g))
			}
			z01.PrintRune('\n')
			continue
		}

		str := string(content)
		for _, j := range str {
			z01.PrintRune(rune(j))
		}
		z01.PrintRune('\n')
	}

}
