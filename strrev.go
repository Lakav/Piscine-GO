package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	var a string
	for i := len(s); i > 0; i-- {
		z01.PrintRune(rune(s[i-1]))
	}
	return a
}
