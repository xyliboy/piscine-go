package piscine

func IsUpper(s string) bool {
	runes := []rune(s)
	length := len(runes)
	i := 0
	j := 0
	for range s {
		if s[i] >= 'A' && s[i] <= 'B' {
			j++
		}
		i++
	}
	if length == j {
		return true
	}
	return false
}
