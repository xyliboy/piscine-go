package piscine

func JumpOver(str string) string {
	word := []rune{}
	runes := []rune(str)
	if len(runes) > 2 {
		for i := 2; i < len(runes); i += 3 {
			word = append(word, runes[i])
		}
	}
	word = append(word, '\n')
	return string(word)
}
