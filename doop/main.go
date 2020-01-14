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

	if l != 3 {
		return
	}
	if a[1] != "+" && a[1] != "-" && a[1] != "/" && a[1] != "*" && a[1] != "%" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if !check(a[0]) || !check(a[2]) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	operator := a[1]

	x := Atoi2(a[0])
	y := Atoi2(a[2])

	if x == -1 && a[0] != "-1" || y == -1 && a[2] != "-1" {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	if overflow(x, y, operator) {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}
	if a[1] == "/" && y == 0 {
		str := "No division by 0"
		for _, str1 := range str {
			z01.PrintRune(str1)
		}
		z01.PrintRune('\n')
		return
	}
	if a[1] == "%" && y == 0 {
		str := "No modulo by 0"
		for _, str1 := range str {
			z01.PrintRune(str1)
		}
		z01.PrintRune('\n')
		return
	}
	c := 0
	switch operator {

	case "+":
		c = x + y
	case "-":
		c = x - y
	case "/":
		c = x / y
	case "*":
		c = x * y
	case "%":
		c = x % y
	}

	c1 := ""
	p := ""
	pl := 0
	if c < 0 {
		c1 = "-"
	}
	if c == 0 {
		z01.PrintRune('0')
	}
	for c != 0 {
		pl = c % 10
		if pl < 0 {
			pl = -pl
		}
		p = string(rune(pl+48)) + p
		c = c / 10
		if c < 0 {
			c = -c
		}
	}
	p = c1 + p

	for _, p2 := range p {
		z01.PrintRune(p2)
	}
	z01.PrintRune('\n')
}

func overflow(x int, y int, o string) bool {
	// max := 9223372036854775807
	// min := -9223372036854775808
	c := 0
	//c1 := 0
	a := false
	// if (x == max && y > 0 && (o == "+" || o == "*")) || (x == min && (y > 0 && (o == "-" || o == "*") || (y < 0 && (o == "+" || o == "*")))) ||
	// 	(y == max && x > 0 && (o == "+" || o == "*")) || (y == min && (x > 0 && o == "*" || (x < 0 && (o == "+" || o == "*")))) {
	// 	a = true
	// }
	if o == "+" {

		c = x + y
		if (x < 0 && y < 0 && c > 0) || (x > 0 && y > 0 && c < 0) {
			a = true
		}

	}
	if o == "-" {
		c = x - y
		if (x < 0 && y > 0 && c > 0) || (x >= 0 && y < 0 && c < 0) {
			a = true
		}
	}
	if o == "*" {
		c = x * y

		// if (x > 0 && y < 0 && c > 0) || (x < 0 && y > 0 && c > 0) || (x > 0 && y > 0 && c < 0) || (x < 0 && y < 0 && c < 0) {
		// 	a = true
		// }
		if x != 0 && c/x != y {
			a = true
		}
	}

	if o == "/" {
		c = x / y
		if (x > 0 && y < 0 && c > 0) || (x < 0 && y > 0 && c > 0) || (x > 0 && y > 0 && c < 0) || (x < 0 && y < 0 && c < 0) {
			a = true
		}
	}

	return a

}

func Atoi2(s string) int {
	c := 1
	p := 0
	l := 0
	for range s {
		l++
	}
	g := []rune(s)
	for i := 0; i < l; i++ {

		if g[i] == '+' {
			continue
		}
		if g[i] == '-' {
			c = -1
			continue
		}
		p = int(g[i]) - 48 + p*10
		//fmt.Println(p)
		if p < 0 && p == -9223372036854775808 && s != "-9223372036854775808" {
			return -1
		}
	}
	//fmt.Println(p)
	return p * c
}

func check(s string) bool {
	count := 0
	for range s {
		count++
	}
	for i, a := range s {
		if !(a >= '0' && a <= '9' || a == '-' && i == 0 && count != 1 || a == '+' && i == 0 && count != 1) {
			return false
		}

	}
	return true
}
