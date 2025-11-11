package piscine

func FindNextPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	i := 2
	for i < nb {
		if nb%i == 0 {
			return false
		}
		i++
	}
	return true
}

func NextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
