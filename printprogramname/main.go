package main

import (
	"os" // για να πάρουμε τα arguments

	"github.com/01-edu/z01" // για εκτύπωση χαρακτήρα-χαρακτήρα
)

func main() {
	name := os.Args[0] // το όνομα του προγράμματος (πρώτο argument πάντα)

	for _, r := range name { // για κάθε χαρακτήρα (rune) στο όνομα
		z01.PrintRune(r) // τύπωσε τον χαρακτήρα
	}
	z01.PrintRune('\n') // αλλαγή γραμμής στο τέλος
}
