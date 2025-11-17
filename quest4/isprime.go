package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	i := 2
	for i <= nb/2 {
		if nb%i == 0 {
			return false
		}
		i++
	}
	return true
}
