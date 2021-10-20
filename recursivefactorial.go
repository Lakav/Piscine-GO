package piscine

func RecursiveFactorial(nb int) int {
	if nb < 26 && nb > 0 {
		nb *= RecursiveFactorial(nb - 1)
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
	return nb
}
