package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n < 0 {
		return
	}

	count := [10]int{}

	// Ειδική περίπτωση για 0
	if n == 0 {
		count[0]++
	}

	// Βάζουμε κάθε ψηφίο στο count
	for n > 0 {
		count[n%10]++
		n /= 10
	}

	// Τυπώνουμε 0 -> 9 με σειρά
	for i := 0; i < 10; i++ {
		for count[i] > 0 {
			z01.PrintRune(rune(i) + '0')
			count[i]--
		}
	}
}
