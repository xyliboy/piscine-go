package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	i := 0
	for range s {
		if runes[i] <= '0' || runes[i] > '0' {
			return true
		}
		i++
	}
	return false
}
