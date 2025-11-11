package piscine

func AlphaCount(s string) int {
	Runes := []rune(s)
	if s == " " {
		return 0
	}
	i := 0
	length1 := 0
	for range s {
		if (Runes[i] >= 'a' && Runes[i] <= 'z') || (Runes[i] <= 'Z' && Runes[i] >= 'A') {
			length1 = length1 + 1
		}
		i++
	}
	return length1
}
