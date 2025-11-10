package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}

	sum := 0
	if nb > 1 && nb <= 20 {
		sum = 1
		for i := 2; i <= nb; i++ {
			sum = sum * i
		}
	}
	if sum < 0 {
		return 0
	}
	return sum
}
