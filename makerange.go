package piscine

func MakeRange(min, max int) []int {
	var nums []int
	dif := max - min
	if dif > 0 {
		nums = make([]int, dif)
		for i := 0; i <= dif-1; i++ {
			nums[i] = min + i
		}
	}
	return nums
}
