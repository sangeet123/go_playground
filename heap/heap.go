package heap

type integers []int

type integersPtr *[]int
type compare func(i int, j int, nos []int) int

func HeapfyInt(nos integers, comp compare) {
	size := len(nos)
	for i := size >> 1; i >= 0; i-- {
		PerformHeapify(nos, i, comp)
	}
}

func PerformHeapify(nos integers, from int, comp compare) {
	size := len(nos)
	loop := true
	for j, l, r := updateIndex(from); l < size && loop; {
		ind := l

		if r < size {
			ind = comp(r, l, nos)
		}

		if comp(ind, j, nos) == ind {
			nos[ind], nos[j] = nos[j], nos[ind]
			j, l, r = updateIndex(ind)
		} else {
			loop = false
		}
	}
}

func Delete(nosPtr integersPtr, comp compare) int {
	if len(*nosPtr) == 0 {
		panic("Empty heap")
	}
	nos := *nosPtr
	toReturn := nos[0]
	nos[0] = nos[len(nos)-1]
	nos = nos[:len(nos)-1]
	PerformHeapify(nos, 0, comp)
	*nosPtr = nos
	return toReturn
}

func updateIndex(start int) (int, int, int) {
	return start, start<<1 + 1, start<<1 + 2
}
