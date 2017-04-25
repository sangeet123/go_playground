package list

func (list List) IsListPalindrum() bool {

	if list.Size() == 0 {
		return false
	}

	if list.Size() == 1 {
		return true
	}

	itForward := list.Iterator()
	itBackward := list.Iterator()

	for i := 0; i < list.Size()-1; i++ {
		itBackward.Next()
	}

	for i := 0; i < list.Size()>>1; i++ {
		if itForward.Next() != itBackward.Prev() {
			return false
		}
	}

	return true

}
