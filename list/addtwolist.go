package list

// no type checking is done
func AddList(l1, l2 *List) *List {

	if l1.Size() == 0 && l2.Size() == 0 {
		return new(List)
	}

	it1 := l1.Iterator()
	it2 := l2.Iterator()

	for i := 0; i < l1.Size()-1; i++ {
		it1.Next()
	}

	for i := 0; i < l2.Size()-1; i++ {
		it2.Next()
	}

	sumList := new(List)
	carry := 0

	// using or condition here and
	// again checking for HasPrev() on
	// a body of for loop is to make sure
	// that we do not have null pointer.
	// We could have done this using other
	// approach which could avoid the use
	// of HasPrev inside the body but its
	// debatable.
	for it1.HasPrev() || it2.HasPrev() {
		no1 := 0
		no2 := 0

		if it1.HasPrev() {
			no1, _ = it1.Prev().(int)
		}

		if it2.HasPrev() {
			no2, _ = it2.Prev().(int)
		}
		sum := no1 + no2 + carry
		newNo := sum

		if sum > 9 {
			carry = 1
			newNo = sum - 10
		} else {
			carry = 0
		}

		sumList.Insert(newNo)
	}

	if carry == 1 {
		sumList.Insert(carry)
	}

	return sumList
}
