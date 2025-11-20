package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	lista := map[string]int{}
	word := ""

	for _, ch := range str {
		if ch == ' ' {
			// ΠΑΝΤΑ καταγράφω την τρέχουσα "λέξη"
			// ακόμα και αν είναι ""
			lista[word]++
			word = ""
		} else {
			word += string(ch)
		}
	}

	// Και μετά το τέλος του string
	// ξανακαταγράφω την τελευταία "λέξη"
	lista[word]++

	return lista
}
