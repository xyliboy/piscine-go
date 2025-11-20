package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	index := 0
	player := 1

	for player <= 4 {
		fmt.Printf("Player %d: %d, %d, %d\n",
			player,
			deck[index],
			deck[index+1],
			deck[index+2],
		)

		index += 3
		player++
	}
}
