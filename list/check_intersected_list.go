package list

import "math/rand"

func GetIntersectedList() (*List, *List, int, int) {

	l1 := new(List)
	l2 := new(List)
	r := rand.New(rand.NewSource(99))
	pos1 := r.Intn(5) + 5
	pos2 := 0
	size := pos1 << 2

	for i := 0; i < size; i++ {
		if i >= (pos1 - 1) {
			node := get_node(r.Intn(100), l1.end, nil)
			// only forward link shares the intersaction
			l1.end.next, l2.end.next = node, node
			l1.end, l2.end = l1.end.next, l2.end.next
			l1.size++
			l2.size++
		} else {
			l1.Append(r.Intn(100))
			for i := 0; i <= r.Intn(2); i++ {
				pos2++
				l2.Append(r.Intn(100))
			}
		}
	}

	return l1, l2, pos1, pos2 + 1
}

func AreListIntersected(l1, l2 *List) (bool, int) {
	s1 := l1.Size()
	s2 := l2.Size()

	forward1 := l1.root
	forward2 := l2.root
	if eq := s1 - s2; eq > 0 {
		for i := 0; i < eq; i++ {
			forward1 = forward1.next
		}
	} else if eq < 0 {
		eq = -eq
		for i := 0; i < eq; i++ {
			forward2 = forward2.next
		}
	}

	for forward1 != nil && forward2 != nil {
		if forward1 == forward2 {
			val, _ := forward1.data.(int)
			return true, val
		}
		forward1, forward2 = forward1.next, forward2.next
	}

	return false, -1
}
