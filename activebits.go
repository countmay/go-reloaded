package student

func ActiveBits(n int) uint {

	var mod int
	count := 1
	if n < 0 {
		n = -n
	}
	for n != 1 {
		mod = n % 2
		n = n / 2
		if mod == 1 {
			count++
		}
	}
	return uint(count)

}
