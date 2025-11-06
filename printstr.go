package piscine

import "github.com/01-edu/z01"

// func PrintStr(a *string) {
// 	*a= *a + "."
// 	b= *a
// 	b := []byte(b)
// 	c=0
// 	for b/="."
// 	z01.PrintRune (byte(b))
// }

func PrintStr(a *string) {
	s := *a
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
}
