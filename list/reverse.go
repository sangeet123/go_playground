package list

func (list *List) Reverse() {
	if list.Size() <= 1 {
		return
	}

	start, newRoot, nextToStart := list.root, list.root, list.root.next
	for start != nil {
		start.prev, start.next = start.next, start.prev
		newRoot = start
		start = nextToStart
		if nextToStart != nil {
			nextToStart = nextToStart.next
		}
	}
	list.root = newRoot
}
