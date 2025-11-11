package piscine

func TrimAtoi(s string) int {
	sign := 1
	n := 0
	found := false

	for _, r := range s {
		if r == '-' && !found {
			sign = -1
		} else if r >= '0' && r <= '9' {
			found = true
			n = n*10 + int(r-'0')
		} else if found {
			return 0
		}
	}

	if !found {
		return 0
	}

	return n * sign
}
