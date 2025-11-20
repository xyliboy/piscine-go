package piscine

func Unmatch(a []int) int {
	count := 0
	met := 0
	for i := range a {
		for j := range a {
			if i != j {
				if a[i] == a[j] {
					count = 0
					met = 0
				} else {
					count++
					met = i
				}
			}
		}
	}
	if count == 0 {
		return -1
	} else {
		return a[met]
	}
}
