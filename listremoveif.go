package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListRemoveIf(l *List, data_ref interface{}) {
	for l.Head != nil && l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}

	// Αν η λίστα άδειασε πλήρως
	if l.Head == nil {
		l.Tail = nil
		return
	}

	// 2) Τώρα αφαιρούμε από το υπόλοιπο της λίστας
	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if curr.Data == data_ref {
			prev.Next = curr.Next
		} else {
			prev = curr
		}
		curr = curr.Next
	}

	// 3) Διόρθωση Tail
	l.Tail = prev
}
