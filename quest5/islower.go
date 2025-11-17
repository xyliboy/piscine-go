package piscine

func IsLower(s string) bool {
	runes := []rune(s)
	length := len(runes)
	i := 0
	j := 0
	for range s {
		if s[i] >= 'a' && s[i] <= 'z' {
			j++
		}
		i++
	}
	if length == j {
		return true
	}
	return false
}
