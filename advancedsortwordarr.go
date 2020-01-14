package student

func AdvancedSortWordArr(array []string, f func(a, b string) int) {
	l := 0
	for range array {
		l++
	}
	k := 0
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			k = f(array[i], array[j])
			if k == 1 {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

func Compare(a, b string) int {
	if b > a {
		return 0
	}
	return 1
}
