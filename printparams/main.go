package printparams

import (
	"os" // για να πάρουμε το path του προγράμματος

	"github.com/01-edu/z01" // για PrintRune
)

func main() {
	args := os.Args[1:]
	for _, line := range args {
		for _, j := range line {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
