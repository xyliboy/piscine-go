package piscine

func LoafOfBread(str string) string {
	// 1. Αφαιρούμε τα spaces
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	// 2. Αν είναι λιγότερο από 5 → λάθος
	if len(clean) < 5 {
		return "Invalid Output\n"
	}

	// 3. Παίρνουμε blocks των 5 + skip 1
	result := ""
	i := 0

	for i+5 <= len(clean) {
		// Πάρε 5 χαρακτήρες
		result += clean[i : i+5]
		result += " "
		i += 6 // προχωράμε 5 + 1 skipped char
	}

	// 4. Αν το αποτέλεσμα τελειώνει με space → κόψτο
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
