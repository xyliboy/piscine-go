package piscine

func AppendRange(min, max int) []int {
	var nums []int
	dif := max - min
	if dif > 0 {
		for i := 0; i <= dif-1; i++ {
			nums = append(nums, min+i)
		}
	}
	return nums
}
