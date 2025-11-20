package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, ch := range runes {
		if ch >= 'a' && ch <= 'z' {
			runes[i] = 'a' + (ch-'a'+14)%26
		} else if ch >= 'A' && ch <= 'Z' {
			runes[i] = 'A' + (ch-'A'+14)%26
		}
	}
	return string(runes)
}
