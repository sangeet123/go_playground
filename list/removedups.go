package list

//using the same list
func (list *List) RemoveDupsInplace() {
	prev := list.root
	next := list.root
	itemMap := make(map[custom_type]struct{})

	for next != nil {
		if _, ok := itemMap[next.data]; !ok {
			itemMap[next.data] = struct{}{}
			prev = next
			next = next.next
		} else {
			prev.next = next.next
			toDelete := next
			next = next.next
			if next != nil {
				next.prev = prev
			}
			toDelete.prev = nil
			toDelete.next = nil
			toDelete = nil
			list.size--
		}
	}
}
