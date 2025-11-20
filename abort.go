package piscine

func Abort(a, b, c, d, e int) int {
	var arr [5]int
	arr[0] = a
	arr[1] = b
	arr[2] = c
	arr[3] = d
	arr[4] = e
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr[2]
}
