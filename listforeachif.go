package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	node := l.Head    // ξεκινάμε από τον πρώτο κόμβο της λίστας
	for node != nil { // όσο δεν έχουμε φτάσει στο τέλος (nil)
		if cond(node) { // αν ο κόμβος περνάει τη συνθήκη cond (επιστρέφει true)
			f(node) // τότε εφαρμόζουμε τη συνάρτηση f πάνω σε αυτόν τον κόμβο
		} // τέλος if
		node = node.Next // προχωράμε στον επόμενο κόμβο της λίστας
	}
}
