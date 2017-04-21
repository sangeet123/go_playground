package list

import "testing"

//ensure all list interface has been implemented
var _ ListInterface = (*List)(nil)

func prepare_list() *List {
	list := new(List)
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	return list
}

// Testing insert method
func TestInsert(t *testing.T) {
	list := prepare_list()
	it := list.Iterator()
	values := []int{4, 3, 2, 1}

	if list.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}

}

// Testing append method
func TestAppend(t *testing.T) {
	list := new(List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	it := list.Iterator()
	values := []int{1, 2, 3, 4}

	if list.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _, val := range values {
		val_in_list := it.Next()
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}

}

// Testing get method
func TestGet(t *testing.T) {
	list := prepare_list()
	values := []int{4, 3, 2, 1}

	i := 0
	for _, val := range values {
		val_in_list := list.Get(i)
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
		i++
	}
}

// Testing delete method
func TestDelete(t *testing.T) {
	list := prepare_list()

	list.Delete(4)
	list.Delete(1)
	list.Delete(2)

	if list.Size() != 1 {
		t.Error("Expected 1 but got ", list.Size())
	}

	if list.Get(0) != 3 {
		t.Error("Expected 3 but got ", list.Get(0))
	}

}

//Testing ISPresent method
func TestIsPresent(t *testing.T) {
	list := prepare_list()
	values := []int{4, 3, 2, 1}
	i := 0
	for _, val := range values {
		val_in_list := list.Get(i)
		if val != val_in_list {
			t.Error("Expected ", val, " but received ", val_in_list)
		}
		i++
	}
}

// Testing InsertAt method
func TestInsertAt(t *testing.T) {
	list := prepare_list()
	list.InsertAt(5, 0)
	list.InsertAt(6, 5)
	list.InsertAt(13, 2)
	values := []int{5, 4, 13, 3, 2, 1, 6}
	for _, val := range values {
		if !list.IsPresent(val) {
			t.Error("Expected true but received false for ", val)
		}
	}
}

//Testing Iterator
func TestIterator(t *testing.T) {
	list := prepare_list()
	it := list.Iterator()

	//move one step forward
	val := it.Next()
	if val != 4 {
		t.Error("Expected 4 but received ", val)
	}

	// move two step backward
	val = it.Prev()
	if val != 3 {
		t.Error("Expected 3 but received ", val)
	}

	val = it.Prev()
	if val != 4 {
		t.Error("Expected 3 but received ", val)
	}

	// check both HasNext and HasPrev returns false
	if it.HasNext() != false {
		t.Error("Expected false but got true")
	}

	if it.HasPrev() != false {
		t.Error("Expected false but got true")
	}

	// reinitialize iterator and move forward till end
	it = list.Iterator()
	it.Next()
	it.Next()
	it.Next()

	values := []int{4, 3, 2, 1}

	i := 3
	for it.HasPrev() {
		val := it.Prev()
		if values[i] != val {
			t.Error("Expected ", values[i], " but received ", val)
		}
		i--
	}
}

//Testing the validaty of lib after some sequence
func TestList(t *testing.T) {
	list := prepare_list()

	list.InsertAt(12, 0)
	list.InsertAt(14, 3)
	list.Append(21)
	list.Insert(100)
	//============ At this point the list must be 100,12,4,3,2,14,1,21=================================//

	// checking the state of list
	//doing some delete and insert operation interchangebly

	values := []int{100, 12, 4, 3, 14, 2, 1, 21}
	assertCorrectListState(list, values, t)

	list.Delete(100) // 12,4,3,14,2,1,21
	values = []int{12, 4, 3, 14, 2, 1, 21}
	assertCorrectListState(list, values, t)

	list.InsertAt(1201, 1) // 12,1201,4,3,14,2,1,21
	values = []int{12, 1201, 4, 3, 14, 2, 1, 21}
	assertCorrectListState(list, values, t)

	list.Delete(12) //1201,4,3,14,2,1,21
	values = []int{1201, 4, 3, 14, 2, 1, 21}
	assertCorrectListState(list, values, t)

	list.Delete(21) //1201,4,3,14,2,1
	values = []int{1201, 4, 3, 14, 2, 1}
	assertCorrectListState(list, values, t)

	list.Append(1202) //1201,4,3,14,2,1
	values = []int{1201, 4, 3, 14, 2, 1, 1202}
	assertCorrectListState(list, values, t)

	list.Delete(2) //1201,4,3,14,1,1202
	values = []int{1201, 4, 3, 14, 1, 1202}
	assertCorrectListState(list, values, t)

	list.InsertAt(1500, 3)
	values = []int{1201, 4, 3, 1500, 14, 1, 1202}
	assertCorrectListState(list, values, t)

	list.InsertAt(1505, 5)
	values = []int{1201, 4, 3, 1500, 14, 1505, 1, 1202}
	assertCorrectListState(list, values, t)

	// Ensure alfter all the operations (Insert, InsertAt, Delete) mixup the list is in consistent state
	assertCorrectBackwardIteration(list, values, t)
}

func assertCorrectBackwardIteration(list *List, values []int, t *testing.T) {
	i := len(values) - 1
	it := list.Iterator()

	for j := i; j > 0; j-- {
		it.Next()
	}

	for it.HasPrev() {
		val := it.Prev()
		if values[i] != val {
			t.Error("Expected ", values[i], " but received ", val)
		}
		i--
	}
}

func assertCorrectListState(list *List, values []int, t *testing.T) {
	it := list.Iterator()
	i := 0
	for it.HasNext() {
		val := it.Next()
		if values[i] != val {
			t.Error("Expected ", values[i], " but received ", val)
		}
		i++
	}
}
