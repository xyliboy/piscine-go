package main

import (
	"os" // για να πάρουμε το path του προγράμματος

	"github.com/01-edu/z01" // για PrintRune
)

func main() {
	path := os.Args[0]    // π.χ. "/home/user/go/choumi" ή "./choumi"
	runes := []rune(path) // μετατροπή σε runes για σωστή επεξεργασία χαρακτήρων

	start := 0                             // index από όπου ξεκινάει το όνομα
	for i := len(runes) - 1; i >= 0; i-- { // πάμε από το τέλος προς τα πίσω
		if runes[i] == '/' { // αν βρούμε "/", άρα εκεί τελειώνει το path
			start = i + 1 // το όνομα ξεκινάει μετά από αυτό
			break
		}
	}

	for j := start; j < len(runes); j++ { // τυπώνουμε από start μέχρι τέλος
		z01.PrintRune(runes[j]) // εκτύπωση κάθε χαρακτήρα
	}
	z01.PrintRune('\n') // αλλαγή γραμμής στο τέλος
}
