package piscine

func DescendAppendRange(max, min int) []int {
	nums := []int{}
	dif := max - min
	if dif > 0 {
		for i := 0; i < dif; i++ {
			nums = append(nums, max-i)
		}
	}
	return nums
}
