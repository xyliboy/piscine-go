package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	node := l.Head    // ξεκινάμε από το πρώτο node
	for node != nil { // όσο υπάρχουν nodes στην λίστα
		if comp(node.Data, ref) { // αν comp λέει ότι είναι ίσα
			return &node.Data // επιστρέφουμε pointer στο Data
		}
		node = node.Next // πάμε στο επόμενο
	}
	return nil
}
