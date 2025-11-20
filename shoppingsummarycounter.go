package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	lista := map[string]int{}
	word := ""

	for _, ch := range str {
		if (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			word += string(ch)
		} else {
			if word != "" {
				lista[word]++
				word = ""
			}
		}
	}

	if word != "" {
		lista[word]++
	}

	return lista
}
