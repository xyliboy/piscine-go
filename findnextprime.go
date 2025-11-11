package piscine

func FindNextPrime(nb int) int {
	if nb < 2 {
		return 2
	}
	for {
		isPrime := true
		if nb%2 == 0 && nb != 2 {
			isPrime = false
		} else {
			for i := 3; i*i <= nb; i += 2 {
				if nb%i == 0 {
					isPrime = false
					break
				}
			}
		}
		if isPrime {
			return nb
		}
		nb++
	}
}
