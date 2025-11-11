package piscine

func IsPrintable(s string) bool {
	runes := []rune(s)
	i := 0
	for range s {
		if runes[i] == '\\' {
			return false
		}
		i++
	}
	return true
}
