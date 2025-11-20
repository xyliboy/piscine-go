package piscine

func StringToIntSlice(str string) []int {
	num := []int{}
	for _, r := range str {
		num = append(num, int(r))
	}
	return num
}
