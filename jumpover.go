package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	if len(str) > 2 {
		runes := []rune(str)
		for i := 2; i <= len(str); i += 3 {
			z01.PrintRune(runes[i])
		}
	}
	return "\n"
}
