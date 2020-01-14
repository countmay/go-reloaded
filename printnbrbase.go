package student

import "github.com/01-edu/z01"

func PrintNbrBase(nbr int, base string) {
	l := 0
	l2 := 0
	c := 0
	p := ""
	j := ""
	for range base {
		l++
	}
	if l < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	a := []rune(base)
	for _, a1 := range a {
		if a1 == '-' || a1 == '+' {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if a[i] == a[j] {
				l2++
			}
		}
		if l2 != 1 {
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		l2 = 0
	}
	

	for nbr != 0 {
		c = nbr % l
		nbr = nbr / l
		if nbr < 0 {
			nbr = -nbr
			j = "-"
		}
		if c<0{
			c=-c
		}
		p = string(a[c]) + p

	}
	p = j + p
	for _, d := range p {
		z01.PrintRune(d)

	}
	z01.PrintRune('\n')

}
