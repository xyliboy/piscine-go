package piscine

func MakeRange(min, max int) []int {
	var nums []int
	dif := max - min
	if dif > 0 {
		for i := 0; i <= dif-1; i++ {
			nums = make([]int, min+i)
		}
	}
	return nums
}
