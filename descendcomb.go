package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for i := 9; i == 1; i-- {
		for j := i - 1; j == 0; j-- {
			z01.PrintRune(rune(i))
			z01.PrintRune(rune(j))
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	for x := 9; x > 0; x-- {
		z01.PrintRune('0')
		z01.PrintRune(rune(x))
	}
	z01.PrintRune('0')
	z01.PrintRune('0')
}
