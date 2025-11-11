package piscine

func TrimAtoi(s string) int {
	result := 0
	foundDigit := false
	negative := false

	for i := 0; i < len(s); i++ {
		// Αν δούμε '-' και δεν έχουμε βρει ψηφίο ακόμα, τότε negative = true
		if s[i] == '-' && !foundDigit {
			negative = true
		}

		// Αν είναι ψηφίο, το προσθέτουμε στον αριθμό
		if s[i] >= '0' && s[i] <= '9' {
			foundDigit = true
			result = result*10 + int(s[i]-'0')
		}
	}

	if !foundDigit {
		return 0
	}

	if negative {
		return -result
	}

	return result
}
