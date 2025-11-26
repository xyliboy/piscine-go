package piscine

func ListSort(l *NodeI) *NodeI {
	if l == nil {
		return nil
	}
	count := 0
	node := l
	for node != nil {
		count++
		node = node.Next
	}

	for i := 0; i < count; i++ {
		curr := l
		for curr.Next != nil {
			if curr.Data > curr.Next.Data {
				curr.Data, curr.Next.Data = curr.Next.Data, curr.Data
			}
			curr = curr.Next
		}
	}

	return l
}
