package piscine

import "github.com/01-edu/z01"

func JumpOver(str string) string {
	runes := []rune(str)
	if len(runes) > 2 {
		for i := 2; i < len(runes); i += 3 {
			z01.PrintRune(runes[i])
		}
	}
	return "\n"
}
