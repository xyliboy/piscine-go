package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	i := 0
	for range s {
		if runes[i] < 32 || runes[i] > 126 {
			return false
		}
		i++
	}
	return true
}
