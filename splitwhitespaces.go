package student

//SplitWhiteSpaces ajjaja
func SplitWhiteSpaces(str string) []string {
	l := 0
	count := 0
	for range str {
		l++
	}
	a := []rune(str)
	if l == 1 {
		b := make([]string, 1)
		if a[0] != '\t' && a[0] != '\n' && a[0] != ' ' {
			b[0] = string(a[0])
		}
		return b
	}
	for i := 1; i < l; i++ {
		if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && i == 1 && (a[0] != '\t' && a[0] != '\n' && a[0] != ' ') {
			count++
		}
		if (a[i] == '\t' || a[i] == '\n' || a[i] == ' ') && i == 1 && (a[0] != '\t' && a[0] != '\n' && a[0] != ' ') {
			count++
		}
		if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && (a[i-1] == '\t' || a[i-1] == '\n' || a[i-1] == ' ') {
			count++
		}

	}

	b := make([]string, count)
	p := ""
	count1 := 0
	for i := 1; i < l; i++ {
		if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && i == 1 && (a[0] != '\t' && a[0] != '\n' && a[0] != ' ') {
			count1++
			p = string(a[0])

		}
		if (a[i] == '\t' || a[i] == '\n' || a[i] == ' ') && i == 1 && (a[0] != '\t' && a[0] != '\n' && a[0] != ' ') {
			count1++
			// p = string(a[0])
			// b[count1-1] = p
			// p = ""
		}
		if (a[i] == '\t' || a[i] == '\n' || a[i] == ' ') && p != "" {
			//|| i == l-1 {
			// if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && i == l-1 {
			// 	p = p + string(a[l-1])
			// }
			b[count1-1] = p
			p = ""

		}
		if i == l-1 && (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') {

			p = p + string(a[i])

			b[count1-1] = p
		}
		if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && (a[i-1] == '\t' || a[i-1] == '\n' || a[i-1] == ' ') {
			count1++
			p = string(a[i])
		}

		if (a[i] != '\t' && a[i] != '\n' && a[i] != ' ') && (a[i-1] != '\t' && a[i-1] != '\n' && a[i-1] != ' ') {
			p = p + string(a[i])
		}

	}

	if (a[1] == '\t' || a[1] == '\n' || a[1] == ' ') && (a[0] != '\t' && a[0] != '\n' && a[0] != ' ') {

		b[0] = string(a[0])
	}
	if (a[l-2] == '\t' || a[l-2] == '\n' || a[l-2] == ' ') && (a[l-1] != '\t' && a[l-1] != '\n' && a[l-1] != ' ') {

		b[count-1] = string(a[l-1])
	}

	return b
}
