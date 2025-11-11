package piscine

func AlphaCount(s string) int {
	Runes := []rune(s)
	length := len(Runes)
	if s == " " {
		return 0
	}
	i := 0
	for range s {
		if Runes[i] == ' ' || Runes[i] >= 0 && Runes[i] <= 9 {
			length = length - 1
			i++
		}
	}
	return length
}
