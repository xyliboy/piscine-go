package piscine

func Max(a []int) int {
	num := 0
	for i := 0; i < len(a); i++ {
		if a[i] >= num {
			num = a[i]
		}
	}
	return num
}
