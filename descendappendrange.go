package piscine

func DescendAppendRange(max, min int) []int {
	nums := make([]int, 0)
	dif := max - min
	if dif > 0 {
		for i := 1; i <= dif; i++ {
			nums = append(nums, min+i)
		}
	}
	return nums
}
