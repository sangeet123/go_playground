package heap

type integers []int

func HeapfyInt(nos integers) {
	size := len(nos)
	heapifyInt(nos, size>>1)
}

func heapifyInt(nos integers, form int) {
	size := len(nos)
	for i := form; i >= 0; i-- {

		j, l, r := updateIndex(i)
		loop := true
		for l < size && loop {
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
}

func updateIndex(start int) (int, int, int) {
	return start, start<<1 + 1, start<<1 + 2
}
