package student

func RecursivePower(nb int, power int) int {
	if power<=0{
      return 0
	}
if power==1{
	return nb
}
nb=nb*RecursivePower(nb, power-1)
return nb
}
