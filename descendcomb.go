package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for a := 98; a > 0; a-- {
		for b := a - 1; b >= 0; b-- {
			z01.PrintRune(rune(a/10 + '0'))
			z01.PrintRune(rune(a%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(b/10 + '0'))
			z01.PrintRune(rune(b%10 + '0'))
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}

	for x := 9; x > 0; x-- {
		z01.PrintRune('0')
		z01.PrintRune(rune(x + '0'))
		z01.PrintRune(',')
		z01.PrintRune(' ')
	}
	z01.PrintRune('0')
	z01.PrintRune('0')
}
