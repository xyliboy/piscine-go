package piscine

func Any(f func(string) bool, a []string) bool {
	num := 0
	for _, i := range a {
		if f(i) == true {
			num = num + 1
		}
	}
	if num == 0 {
		return false
	}
	return true
}
