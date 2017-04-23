package list

//using the same list
func (list *List) RemoveDupsInplace() {
	prev := list.root
	next := list.root
	itemMap := make(map[custom_type]bool)

	for next != nil {
		if _, ok := itemMap[next.data]; !ok {
			itemMap[next.data] = true
			prev = next
			next = next.next
		} else {
			prev.next = next.next
			toDelete := next
			next = next.next
			toDelete.prev = nil
			toDelete.next = nil
			list.size--
		}
	}
}
