package main

import (
	"os" // για να πάρουμε το path του προγράμματος

	"github.com/01-edu/z01" // για PrintRune
)

func main() {
	args := os.Args[1:]
	for i := len(args) - 1; i >= 0; i-- {
		for _, j := range args[i] {
			z01.PrintRune(j)
		}

		z01.PrintRune('\n')
	}
}
