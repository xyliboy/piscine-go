package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	lista := map[string]int{}
	word := ""
	for _, ch := range str { // περνάω χαρακτήρα-χαρακτήρα το string
		if ch != ' ' { // όσο δεν είναι space, συνεχίζω το χτίσιμο
			word = word + string(ch)
		} else { // όταν βρω space, η λέξη τελείωσε
			if word != "" { // αν όντως υπάρχει λέξη
				lista[word]++
				word = "" // καθαρίζω τη λέξη για την επόμενη
			}
		}
	}
	if word != "" { // καταγράφω την τελευταία λέξη
		lista[word]++
	}
	return lista
}
