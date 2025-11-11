package piscine

func Index(s string, toFind string) int {
	runess := []rune(s)
	runest := []rune(toFind)
	lengtht := len(toFind)
	i := 0
	j := 0
	for range s {
		if runess[i] == runest[0] {
			for range toFind {
				a := i
				if runess[a] == runest[j] {
					j++
					a++
				}
				return 0
			}
			if j == lengtht {
				return i
			}
			j = 0
		}
		i++
	}
	return 0
}
