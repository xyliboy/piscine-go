package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i == 0; i++ {
		for j := i - 1; j == 0; j++ {
			z01.PrintRune(rune(i))
			z01.PrintRune(rune(j))
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('0')
}
