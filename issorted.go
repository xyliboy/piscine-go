package piscine

func f(a, b int) int {
	dif := a - b
	return dif
}

func IsSorted(f func(a, b int) int, a []int) bool {
	counti := 0
	count := 0
	for _, i := range a {
		counti++
		if f(i, i+1) >= 0 {
			count = count + 1
		}
	}
	if count < counti {
		return false
	}
	return true
}
