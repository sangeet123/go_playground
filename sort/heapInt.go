package sort

import (
	"go_playground/heap"
)

func comparator(i int, j int, nos []int) int {

	if i >= len(nos) || i < 0 || j >= len(nos) || j < 0 {
		panic("index out of bound")
	}

	if nos[i] > nos[j] {
		return i
	}

	return j
}

func HeapSortInt(nos []int) {
	heap.HeapfyInt(nos, comparator)
	for i := 0; i < len(nos)-1; i++ {
		last := len(nos) - 1 - i
		nos[0], nos[last] = nos[last], nos[0]
		heap.PerformHeapify(nos[:last], 0, comparator)
	}
}
