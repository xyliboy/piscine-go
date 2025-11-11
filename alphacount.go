package piscine

func AlphaCount(s string) int {
	Runes := []rune(s)
	length := len(Runes)
	if s == " " {
		return 0
	}
	i := 0
	for range s {
		if Runes[i] == ' ' {
			length = length - 1
			i++
		}
	}
	return length
}
