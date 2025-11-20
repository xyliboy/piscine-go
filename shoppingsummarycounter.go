package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	lista := map[string]int{}
	word := ""

	for _, ch := range str {
		if ch != ' ' && ch != '\t' && ch != '\n' && ch != '\r' {
			word = word + string(ch)
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
