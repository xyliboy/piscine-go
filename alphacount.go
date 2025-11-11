package piscine

func AlphaCount(s string) int {
	Runes := []rune(s)
	length := len(Runes)
	if s == " " {
		return 0
	}
	return length
}
