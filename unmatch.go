// package piscine

// func Unmatch(a []int) int {
// 	count := 0
// 	met := 0
// 	for i := range a {
// 		for j := range a {
// 			if i != j {
// 				if a[i] == a[j] {
// 					count = 0
// 				} else {
// 					count++
// 				}
// 			}
// 		}
// 		if count != 0 {
// 			met = i
// 			return a[met]
// 		}

// 	}

// 	return -1
// }

package piscine

func Unmatch(a []int) int {
	for i, num := range a {
		count := 0
		for j := range a {
			if i != j {
				if a[i] == a[j] {
					count++
				}
			}
		}
		if count%2 != 0 {
			return num
		}

	}

	return -1
}
