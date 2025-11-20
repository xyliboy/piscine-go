package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	if order == "burger" {
		burger := food{
			preptime: 15,
		}
		return burger.preptime
	}
	if order == "chips" {
		chips := food{
			preptime: 10,
		}
		return chips.preptime
	}
	if order == "nuggets" {
		nuggets := food{
			preptime: 12,
		}
		return nuggets.preptime
	}
	return 404
}
