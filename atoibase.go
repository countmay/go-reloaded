package student

func AtoiBase(s string, base string) int {

	l := 0
	l2 := 0
	l3 := 0
	p := 0

	for range base {
		l++
	}
	if l < 2 {
		return 0
	}
	a := []rune(base)
	for _, a1 := range a {
		if a1 == '-' || a1 == '+' {
			return 0
		}
	}
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			if a[i] == a[j] {
				l2++
			}
		}
		if l2 != 1 {
			return 0
		}
		l2 = 0
	}
	for range s {
		l3++
	}
	s1 := []rune(s)
	for o := 0; o < l3; o++ {
		for k := 0; k < l; k++ {
			if s1[o] == a[k] {
				p = k + p*l

			}
		}
	}

	return p

}
