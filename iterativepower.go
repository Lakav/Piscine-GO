package piscine

func IterativePower(nb int, power int) int {
	start := 1
	for a := 1; a <= power; a++ {
		start *= nb
	}
	return start
}
