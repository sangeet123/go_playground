package list

import "testing"

//ensure all list interface has been implemented
var _ ListInterface = (*List)(nil)

func prepare_list()*List{
	list := new(List)
	list.Insert(1)
	list.Insert(2)
	list.Insert(3)
	list.Insert(4)
	return list
}

// Testing insert method
func TestInsert(t *testing.T){
	list := prepare_list()
	it := list.Iterator()
	values := []int{4,3,2,1}

	if list.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _,val := range values{
		val_in_list := it.Next()
		if val != val_in_list{
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}

}

// Testing append method
func TestAppend(t *testing.T){
	list := new(List)
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)

	it := list.Iterator()
	values := []int{1,2,3,4}

	if list.Size() != 4 {
		t.Error("Expected 4 but got ", list.Size())
	}

	for _,val := range values{
		val_in_list := it.Next()
		if val != val_in_list{
			t.Error("Expected ", val, " but received ", val_in_list)
		}
	}

}

// Testing get method
func TestGet(t *testing.T){
	list := prepare_list()
	values := []int{4,3,2,1}

	i := 0
	for _, val := range values{
		val_in_list := list.Get(i)
		if val != val_in_list{
			t.Error("Expected ", val, " but received ", val_in_list)
		}
		i++		
	}
}

// Testing delete method
func TestDelete(t *testing.T){
	list := prepare_list()

	list.Delete(4)
	list.Delete(1)
	list.Delete(2)

	if list.Size() != 1{
		t.Error("Expected 1 but got ", list.Size())
	}

	if list.Get(0) != 3{
		t.Error("Expected 3 but got ", list.Get(0))
	}

}

func TestIsPresent(t *testing.T){
	list := prepare_list()
	values := []int{4,3,2,1}
	i := 0
	for _, val := range values{
		val_in_list := list.Get(i)
		if val != val_in_list{
			t.Error("Expected ", val, " but received ", val_in_list)
		}
		i++
	}
}


func TestInsertAt(t *testing.T){
	list := prepare_list()
	list.InsertAt(5,0)
	list.InsertAt(6,5)
	list.InsertAt(13,2)
	values := []int{5,4,13,3,2,1,6}
	for _, val := range values{
		if !list.IsPresent(val){
			t.Error("Expected true but received false for ", val)
		}
	}
}

func TestIterator(t *testing.T){
	list := prepare_list()
	it := list.Iterator()

	//move one step forward
	val := it.Next()
	if val != 4{
		t.Error("Expected 4 but received ", val)
	}

	// move two step backward
	val = it.Prev()
	if val != 3{
		t.Error("Expected 3 but received ", val)
	}

	val = it.Prev()
	if val != 4{
		t.Error("Expected 3 but received ", val)
	}

	// check both HasNext and HasPrev returns false
	if it.HasNext() != false{
		t.Error("Expected false but got true")
	}

	if it.HasPrev() != false{
		t.Error("Expected false but got true")
	}

	// reinitialize iterator and move forward till end
	it = list.Iterator()
	it.Next()
	it.Next()
	it.Next()

	values := []int{4,3,2,1}

	i := 3
	for it.HasPrev(){
		val := it.Prev()
		if values[i] != val{
			t.Error("Expected ", values[i], " but received ", val)
		}
		i--
	}
}



