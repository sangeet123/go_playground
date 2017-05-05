package heap

type integers []int

func HeapfyInt(nos integers) {
	size := len(nos)
	for i := size >> 1; i >= 0; i-- {
		heapifyInt(nos, i)
	}
}

func heapifyInt(nos integers, from int) {
	size := len(nos)
	loop := true
	for j, l, r := updateIndex(from); l < size && loop; {
		ind := l

		if r < size && nos[r] < nos[l] {
			ind = r
		}

		if nos[ind] < nos[j] {
			nos[ind], nos[j] = nos[j], nos[ind]
			j, l, r = updateIndex(ind)
		} else {
			loop = false
		}
	}
}

func updateIndex(start int) (int, int, int) {
	return start, start<<1 + 1, start<<1 + 2
}
