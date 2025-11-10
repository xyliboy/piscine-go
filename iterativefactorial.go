package piscine

func IterativeFactorial(nb int) int {
	if nb <= 0 {
		return 0
	}

	sum := 1
	if nb > 1 {
		for i := 2; i <= nb; i++ {
			sum = sum * i
		}
	}

	return sum
}
