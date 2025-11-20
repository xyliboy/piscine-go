package piscine

func Unmatch(a []int) int {
	count := 0
	met := ' '
	for i := range len(a) {
		for j := range len(a) {
			if i != j {
				if a[i] == a[j] {
					count = 0
					met = ' '
				} else {
					count++
					met = rune(a[i] + '0')
				}
			}
		}
	}
	if met == ' ' {
		return -1
	} else {
		return int(met)
	}
}
