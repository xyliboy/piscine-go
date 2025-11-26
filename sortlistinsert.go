package piscine

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	newNode := &NodeI{Data: data_ref} // φτιάχνουμε νέο node

	// 1) Αν η λίστα είναι άδεια → το νέο node γίνεται head
	if l == nil {
		return newNode
	}

	// 2) Αν το νέο στοιχείο είναι μικρότερο από το πρώτο → νέο head
	if data_ref < l.Data {
		newNode.Next = l
		return newNode
	}

	// 3) Αλλιώς βρες τη σωστή θέση στη μέση ή στο τέλος
	curr := l
	for curr.Next != nil && curr.Next.Data < data_ref {
		curr = curr.Next
	}

	// 4) Εισαγωγή του νέου node
	newNode.Next = curr.Next
	curr.Next = newNode

	return l
}
