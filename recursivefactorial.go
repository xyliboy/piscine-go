package piscine

func RecursiveFactorial(nb int) int {
	// λάθος είσοδοι
	if nb < 0 || nb > 20 {
		return 0
	}
	// base case
	if nb == 0 || nb == 1 {
		return 1
	}

	// πρώτα παίρνουμε το factorial του nb-1
	sub := RecursiveFactorial(nb - 1)
	// αν το sub είναι 0 σημαίνει error από κάτω (overflow ή λάθος)
	if sub == 0 {
		return 0
	}

	product := nb * sub
	// έλεγχος overflow με διαίρεση — πιο αξιόπιστος
	if product/sub != nb { // ή product/nb != sub
		return 0
	}

	return product
}
