package student

import (
	"github.com/01-edu/z01"
)

func PrintCombN(n int) {
	PrintCombNew(0, n, "")
	z01.PrintRune('\n')
}

func PrintCombNew(start, n int, out string) {

	if n == 0 {
		if out == "9" || out == "89" || out == "789" || out == "6789" || out == "56789" || out == "456789" || out == "3456789" || out == "23456789" || out == "123456789" {
			for _, j := range out {
				z01.PrintRune(j)
			}
		} else {
			for _, j := range out {
				z01.PrintRune(j)
			}
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	for i := start; i <= 9; i++ {
		str := out + string(i+48)
		PrintCombNew(i+1, n-1, str)
	}
}

// package student

// import "github.com/01-edu/z01"

// //PrintCombN 44
// func PrintCombN(n int) {
// 	if n <= 0 {
// 		return
// 	}

// 	//k1 := 0
// 	p := ""
// 	//k2 := 0
// 	// for i := 0; i < n; i++ {
// 	// 	z01.PrintRune(rune(i + 48))
// 	// 	k = k + 1
// 	// }

// 	//j := 9 - n - 1
// 	for i1 := n; i1 >= 1; i1-- {
// 		k := 0
// 		for i2 := 0; i2 <= (9 - i1); i2++ {
// 			for i := 0; i < i1-1; i++ {
// 				z01.PrintRune(rune(i + 48))
// 				k = k + 1
// 			}
// 			z01.PrintRune(rune(k + i2 + 48))
// 		}
// 		for _, a := range p {
// 			z01.PrintRune(a)
// 		}
// 		p = p + string(rune(i1+9-n))
// 	}
// 	// k = k - 1
// 	// for k != 9 {
// 	// 	for i := 0; i < n-1; i++ {
// 	// 		z01.PrintRune(',')
// 	// 		z01.PrintRune(' ')
// 	// 		z01.PrintRune(rune(i + 48))

// 	// 		k1 = k1 + 1
// 	// 	}
// 	// 	k = k + 1

// 	// 	z01.PrintRune(rune(k + 48))
// 	// 	p = p + "k"
// 	// 	//k2 = k
// 	// }
// 	// n = n - 1
// 	// PrintCombN(n - 1)
// 	// for n != 0 {

// 	// 	k := 0
// 	// 	k1 := 0
// 	// 	for i := 0; i < n; i++ {
// 	// 		z01.PrintRune(rune(i + 48))
// 	// 		k = k + 1
// 	// 	}
// 	// 	for k != k2-1 {
// 	// 		for i := 0; i < n-1; i++ {
// 	// 			z01.PrintRune(rune(i + 48))
// 	// 			k1 = k1 + 1
// 	// 		}
// 	// 		k = k + 1
// 	// 		z01.PrintRune(rune(k + 48))
// 	// 		p = p + "k"
// 	// 		k2 = k
// 	// 	}
// 	// 	n = n - 1

// 	// }

// }
