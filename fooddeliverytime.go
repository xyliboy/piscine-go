package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	faghto := food{}
	fai := 0
	count := 0
	for i := range order {
		if order[i] != ' ' {
			count++
		}
		if order[i] == ' ' {
			if count == 6 {
				fai = fai + 15
			}
			if count == 5 {
				fai = fai + 10
			}
			if count == 7 {
				fai = fai + 12
			}
			count = 0
		}

	}
	faghto.preptime = fai
	return faghto.preptime
}
