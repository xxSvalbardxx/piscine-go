package piscine

func ListSize(l *List) int {
	n := &NodeL{}
	count := 0
	if l.Head == nil {
		l.Head = n
		count++
	} else {
		current := l.Head
		for current.Next != nil {
			current = current.Next
			count++
		}
	}
	return count
}
