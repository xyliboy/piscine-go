package piscine

func RecursiveFactorial(nb int) int {
	sum := 1
	if nb < 0 || nb > 20 {
		sum = 0
	}
	if nb == 0 || nb == 1 {
		sum = 1
	}
	if nb > 1 && nb < 21 {
		sum = nb * RecursiveFactorial(nb-1)
		if sum < 0 {
			return 0
		}
	}
	return sum

}
