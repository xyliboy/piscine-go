package piscine

func ToUpper(s string) string {
	runes := []rune(s)
	if s != "" {
		i := 0
		for range s {
			if runes[i] >= 'a' || runes[i] <= 'z' {
				runes[i] = runes[i] + 64
			}
		}
	}
	return s
}
