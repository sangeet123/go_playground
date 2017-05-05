package sort

func SelectionSortInt(nos integers) {
	for i := 0; i < len(nos)-1; i++ {
		ind := getSmallestElementIndex(i, nos)
		nos[i], nos[ind] = nos[ind], nos[i]
	}
}

func getSmallestElementIndex(from int, nos integers) int {
	smallInd := from

	for i := from; i < len(nos); i++ {
		if nos[smallInd] > nos[i] {
			smallInd = i
		}
	}

	return smallInd
}
