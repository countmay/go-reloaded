package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {

	a := os.Args[1:]
	l := 0

	for range a {
		l++
	}
	if l < 1 {
		z01.PrintRune('\n')
		return
	}
	p := a[0]

	vow := []string{"A", "E", "I", "O", "U", "a", "e", "i", "o", "u"}

	for i := 1; i < l; i++ {
		p = p + " " + a[i]

	}

	l2 := 0
	d := []rune(p)
	for range d {
		l2++
	}
	c := l2 - 1
	n := 0
	flag := false
	for m := 0; m < l2; m++ {
		for _, vow1 := range vow {

			if string(d[m]) == vow1 {
				n = m
				flag = true
			}

			if c > m && string(d[c]) == vow1 && flag == true {
				d[n], d[c] = d[c], d[n]

			}

		}
		c--
	}

	for _, pr := range d {
		z01.PrintRune(pr)
	}
	z01.PrintRune('\n')

}
