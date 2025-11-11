package piscine

func NRune(s string, n int) rune {
	i := -1
	runes := []rune(s)
	for range s {
		i++
	}
	if n == i && i >= 0 {
		return runes[n]
	}
	return 0
}
