package piscine

func IterativeFactorial(nb int) int {
	if nb > 0 {
		nb = 1
		if nb > 1 {
			for i := 2; i <= nb; i++ {
				nb = nb * i
			}
		}
	}
	return nb
}
