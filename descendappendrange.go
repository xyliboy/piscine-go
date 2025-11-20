package piscine

func DescendAppendRange(max, min int) []int {
	var nums []int
	dif := max - min
	if dif > 0 {
		for i := 0; i < dif; i++ {
			nums = append(nums, min+i)
		}
	}
	return nums
}
