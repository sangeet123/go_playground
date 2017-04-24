package list

func (list *List) DeleteMiddle() {
	//return if list is empty
	if list.Size() == 0 {
		return
	}

	if list.Size() <= 2 {
		firstNode := list.root
		list.root = list.root.next
		if list.root != nil {
			list.root.prev = nil
			firstNode.next = nil
		}
		list.size--
		return
	}

	fast := list.root
	slow := fast

	for fast != nil {
		fast = fast.next
		if fast != nil {
			slow = slow.next
			fast = fast.next
		}
	}

	slow.prev.next = slow.next
	slow.next.prev = slow.prev
	slow.prev = nil
	slow.next = nil
	slow = nil
	list.size--

}
