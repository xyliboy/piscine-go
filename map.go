package piscine

func Map(f func(int) bool, a []int) []bool {
	nums := make([]bool, 0)
	for _, i := range a {
		if f(i) == true {
			nums = append(nums, true)
		} else {
			nums = append(nums, false)
		}
	}
	return nums
}
