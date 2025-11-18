package piscine

func first(a, b int) int {
	dif := a - b
	return dif
}

func IsSorted(f func(a, b int) int, a []int) bool {
	counti := 0
	countm := 0
	countp := 0
	for i := 0; i < len(a)-1; i++ {
		counti++
		if f(a[i], a[i+1]) >= 0 {
			countp = countp + 1
		}
		if f(a[i], a[i+1]) <= 0 {
			countm++
		}

	}
	if countp < counti || countm < counti {
		return false
	}
	return true
}
