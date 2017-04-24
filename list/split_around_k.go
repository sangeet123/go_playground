package list

// only works for list of integer
func (list *List) SplitAroundK(k int) {
	if list.Size() <= 1 {
		return
	}

	curPos := list.root
	start := curPos

	for start != nil {
		if data, ok := start.data.(int); ok && data < k {
			curPos.data, start.data = start.data, curPos.data
			curPos = curPos.next
		}
		start = start.next
	}

}
