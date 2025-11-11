package piscine

func Nrune(s string, n int) rune {
	i := -1
	runes := []rune(s)
	for range s {
		i++
	}
	if n <= i-1 {
		return runes[n]
	}
	return 0
}
