package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func DealAPackOfCards(deck []int) {
	// 4 παίκτες, 3 κάρτες ο καθένας
	for player := 1; player <= 4; player++ {

		// fmt για να γράψουμε "Player X:"
		fmt.Print("Player ", player, ":")

		// Από που έως που παίρνει κάρτες ο παίκτης
		start := (player - 1) * 3
		end := start + 3

		fmt.Print(" [") // αρχή λίστας καρτών

		// Μοίρασμα 3 καρτών ανά παίκτη
		for i := start; i < end; i++ {

			// Εκτύπωση αριθμού με fmt
			fmt.Print(deck[i])

			// Αν δεν είναι η τελευταία κάρτα, βάζουμε κόμμα
			if i < end-1 {
				fmt.Print(", ")
			}
		}

		fmt.Print("]")

		// Νέα γραμμή με z01 (για να χρησιμοποιήσουμε και z01 όπως θες)
		z01.PrintRune('\n')
	}
}
