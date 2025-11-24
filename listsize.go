package piscine

func ListSize(l *List) int {
	num := 0
	if l.Tail == nil {
		return num
	}
	for l.Head != nil {
		l.Head = l.Head.Next
		num++
	}
	return num
}
