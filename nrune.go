package piscine

func NRune(s string, n int) rune {
	runes := []rune(s)
	length := len(runes)
	if length < n || n <= 0 {
		return 0
	}
	m := n - 1
	return runes[m]
}
