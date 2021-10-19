package piscine

func StrRev(s string) string {
	var a string
	for i := len(s); i > 0; i-- {
		a += string(s[i-1])
	}
	return a
}
