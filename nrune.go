package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	length := len(runes)
	if length < n {
		return 0
	}
	return runes[n-1]
}
