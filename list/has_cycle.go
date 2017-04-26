package list

import "math/rand"

func CreateListWithCycle() (*List, custom_type) {
	l1 := new(List)
	r := rand.New(rand.NewSource(10000))
	size := 1000000

	for i := 0; i < size; i++ {
		l1.Append(r.Intn(100))
	}

	l1.end.next = l1.root.next.next.next

	return l1, l1.end.next.data.(int)
}

func (l1 List) HasCycle() (bool, custom_type) {

	if l1.Size() <= 1 {
		return false, nil
	}

	fast := l1.root
	slow := fast

	for fast != nil && slow != nil {
		if fast == slow {
			return true, fast.data.(int)
		}
		fast = fast.next
		if fast != nil {
			fast = fast.next
		}

		slow = slow.next
	}

	return false, nil

}
