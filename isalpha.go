package piscine

func IsAlpha(s string) bool {
	runes := []rune(s)
	length := len(runes)
	i := 0
	j := 0
	for range s {
		if (s[i] >= 'a' && s[i] <= 'z' || s[i] >= 'A' && s[i] <= 'Z') || s[i] <= '0' || s[i] > '0' {
			j++
		}
		i++
	}
	if length == j {
		return true
	}
	return false
}
