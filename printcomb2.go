package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 0; a <= 98; a++; {
		for b :=a+1; b <= 99; b++; {
			z01.PrintRune(a + '0')
			z01.PrintRune('')
			Z01.PrintRune(b + '0')
			if a =/ {
				Z01.PrintRune(',')
			}

		}
	}
	Z01.PrintRune('\n')
}
