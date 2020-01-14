package student

func ConvertBase(nbr, baseFrom, baseTo string) string {
	l := 0
	l1 := 0
	l2 := 0
	p := 0
	for range nbr {
		l++
	}
	for range baseFrom {
		l1++
	}
	for range baseTo {
		l2++
	}
	for _, a := range nbr {
		for i, b := range baseFrom {
			if a == b {
				p = i + p*l1
			}
		}
	}
	s := ""
	for p != 0 {
		s = string(baseTo[p%l2]) + s
		p = p / l2
	}
	return s
}
