package piscine

// Ελέγχει αν ένας αριθμός είναι πρώτος
func FindNextPrime(n int) bool {
	if n <= 1 {
		return false
	}
	i := 2
	for i < n {
		if n%i == 0 {
			return false
		}
		i++
	}
	return true
}

// Βρίσκει τον πρώτο πρώτο αριθμό ≥ nb
func NextPrime(nb int) int {
	for {
		if IsPrime(nb) {
			return nb
		}
		nb++
	}
}
