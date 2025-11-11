package piscine

func Index(s string, toFind string) int {
	runess := []rune(s)
	runest := []rune(toFind)
	lengtht := len(runest)
	i := 0
	j := 0
	for range s {
		if runess[i] == runest[0] {
			a := i
			for range toFind {
				if runess[a] == runest[j] {
					j++
					a++
				}
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
