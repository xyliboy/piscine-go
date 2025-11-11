package piscine

func NRune(s string, n int) rune {
	i := -1
	runes := []rune(s)
	for range s {
		i++
	}
	if runes[i-1] == runes[n] {
		return runes[n]
	}
	return 0
}
