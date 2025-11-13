package piscine

func SplitWhiteSpaces(s string) []string {
	var runes []rune
	var pollesleksis []string
	for i, c := range s {
		if c != '\n' && c != '\t' && c != ' ' {
			runes = append(runes, c)
		}
		if (c == ' ' || c == '\n' || c == '\t' || i == len(s)-1) && len(runes) != 0 {
			pollesleksis = append(pollesleksis, string(runes))
			runes = []rune{}
		}
	}
	return pollesleksis
}
