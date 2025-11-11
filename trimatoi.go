package piscine

func TrimAtoi(s string) int {
	result := 0
	found := false

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			found = true
			result = result*10 + int(s[i]-'0')
		}
	}

	if !found {
		return 0
	}

	return result
}
