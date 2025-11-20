package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	var arr [12]int
	count := 0
	for i := 0; i < 12; i += 3 {
		count++
		fmt.Println("Player")
		z01.PrintRune(rune(count))
		z01.PrintRune(':')
		for j := i; j <= i+2; j++ {
			z01.PrintRune(' ')
			z01.PrintRune(rune(arr[j]))
			if (j+1)%3 != 0 {
				z01.PrintRune(',')
			}

		}
		if i != 9 {
			z01.PrintRune('\n')
		}
	}
}
