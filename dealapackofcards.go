package piscine

import "github.com/01-edu/z01"

func printString(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	digits := []rune{}
	for n > 0 {
		digits = append(digits, rune(n%10+'0'))
		n /= 10
	}
	for i := len(digits) - 1; i >= 0; i-- {
		z01.PrintRune(digits[i])
	}
}

func DealAPackOfCards(deck []int) {
	player := 1

	for i := 0; i < 12; i += 3 {
		// Player X:
		printString("Player ")
		printInt(player)
		printString(": [")

		// Τα 3 χαρτιά
		printInt(deck[i])
		printString(", ")
		printInt(deck[i+1])
		printString(", ")
		printInt(deck[i+2])

		printString("]\n")
		player++
	}
}
