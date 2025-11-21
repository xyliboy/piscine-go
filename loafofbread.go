package piscine

func LoafOfBread(str string) string {
	// Αφαιρούμε τα spaces
	clean := ""
	for _, r := range str {
		if r != ' ' {
			clean += string(r)
		}
	}

	// Αν < 5 => error
	if len(clean) < 5 {
		return "Invalid Output\n"
	}

	result := ""
	i := 0

	// Παίρνουμε blocks των 5 και skip 1
	for i+5 <= len(clean) {
		result += clean[i:i+5] + " "
		i += 6
	}

	// Βάζουμε ό,τι περισσεύει
	if i < len(clean) {
		result += clean[i:]
	}

	// ❗ ΑΦΑΙΡΟΥΜΕ space στο τέλος αν υπάρχει
	if len(result) > 0 && result[len(result)-1] == ' ' {
		result = result[:len(result)-1]
	}

	return result + "\n"
}
