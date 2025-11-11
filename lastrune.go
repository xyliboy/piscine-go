package piscine

func LastRune(s string) rune {
	runes := []rune(s)
	i := -1
	for range s {
		i++
	}
	return runes[i]
}
