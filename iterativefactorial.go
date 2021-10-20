package piscine

func IterativeFactorial(nb int) int {
	if nb < 26 && nb > 0 {
		for a := nb - 1; a > 1; a-- {
			nb *= a
		}
	} else if nb == 0 {
		return 1
	} else {
		return 0
	}
	return nb
}
