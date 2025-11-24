package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	var prev *NodeL = nil
	current := l.Head
	var next *NodeL = nil

	if current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev
}
