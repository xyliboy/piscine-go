package piscine

func IsNumeric(s string) bool {
	runes := []rune(s)
	i := 0
	for range s {
		if runes[i] < '0' || runes[i] > '9' {
			return false
		}
		i++
	}
	return true
}
