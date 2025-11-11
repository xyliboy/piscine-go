package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		// Κάνουμε όλα μικρά αρχικά
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = runes[i] + 32
		}

		// Αν είναι αρχή λέξης και γράμμα, κάντο κεφαλαίο
		if i == 0 || !isAlnum(runes[i-1]) {
			if runes[i] >= 'a' && runes[i] <= 'z' {
				runes[i] = runes[i] - 32
			}
		}
	}
	return string(runes)
}

// Βοηθητική: ελέγχει αν είναι γράμμα ή αριθμός
func isAlnum(r rune) bool {
	return (r >= 'A' && r <= 'Z') ||
		(r >= 'a' && r <= 'z') ||
		(r >= '0' && r <= '9')
}
