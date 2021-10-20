package piscine

func IterativeFactorial(nb int) int {
	if nb < 24 && nb > 0 {
		for a := nb - 1; a > 1; a-- {
			nb *= a
		}
	} else {
		return 0
	}
	return nb
}
