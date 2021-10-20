package piscine

func IterativePower(nb int, power int) int {
	start := 1
	if power >= 0 {
		for a := 1; a <= power; a++ {
			start *= nb
		}
	} else {
		return 0
	}
	return start
}
