package list

func (list *List) KthToLast(k int) custom_type {

	if k < 0 || k > list.Size() {
		panic("Cannot access data pass the list size")
	}

	start := list.root

	for i := 0; i < list.Size()-k-1; i++ {
		start = start.next
	}

	return start.data

}
