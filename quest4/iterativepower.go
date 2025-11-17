package piscine

func IterativePower(nb int, power int) int {
	res := nb
	if power < 0 {
		return 0
	}

	if power == 0 {
		return 1
	}

	if power == 1 {
		return nb
	}

	if power > 1 {
		for i := 2; i <= power; i++ {
			res = res * nb
		}
	}
	return res
}
