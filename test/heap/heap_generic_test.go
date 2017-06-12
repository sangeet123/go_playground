package heaptest

import (
	"go_playground/heap"
	"testing"
)

func TestGenericHeapTest(t *testing.T) {
	nos := []heap.Node{heap.Node{Data: 5, Priority: 1},
		heap.Node{Data: 4, Priority: 2},
		heap.Node{Data: 3, Priority: 3},
		heap.Node{Data: 2, Priority: 4},
		heap.Node{Data: 1, Priority: 5}}

	h := heap.NewHeap(nos)

	expectedSequence := []heap.Node{heap.Node{Data: 1, Priority: 5},
		heap.Node{Data: 2, Priority: 4},
		heap.Node{Data: 3, Priority: 3},
		heap.Node{Data: 4, Priority: 2},
		heap.Node{Data: 5, Priority: 1}}

	h.Heapfy()
	for _, val := range expectedSequence {
		deleted := h.Delete()
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}

func TestGenericHeapTestSingleElementArray(t *testing.T) {
	nos := []heap.Node{heap.Node{Data: 1, Priority: 1}}
	expectedSequence := []heap.Node{heap.Node{Data: 1, Priority: 1}}
	h := heap.NewHeap(nos)
	h.Heapfy()
	for _, val := range expectedSequence {
		deleted := h.Delete()
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}

func TestGenericHeapTestTwoElementsArray(t *testing.T) {
	nos := []heap.Node{heap.Node{Data: 2, Priority: 1}, heap.Node{Data: 1, Priority: 2}}
	expectedSequence := []heap.Node{heap.Node{Data: 1, Priority: 2}, heap.Node{Data: 2, Priority: 1}}
	h := heap.NewHeap(nos)
	h.Heapfy()
	for _, val := range expectedSequence {
		deleted := h.Delete()
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}

func TestGenericHeapTestExample(t *testing.T) {
	nos := []heap.Node{
		heap.Node{Data: 15, Priority: 2},
		heap.Node{Data: 17, Priority: 1},
		heap.Node{Data: 2, Priority: 10},
		heap.Node{Data: 4, Priority: 8},
		heap.Node{Data: 12, Priority: 4},
		heap.Node{Data: 15, Priority: 2},
		heap.Node{Data: 17, Priority: 1},
		heap.Node{Data: 8, Priority: 6},
		heap.Node{Data: 2, Priority: 10}}

	expectedSequence := []heap.Node{
		heap.Node{Data: 2, Priority: 10},
		heap.Node{Data: 4, Priority: 8},
		heap.Node{Data: 8, Priority: 6},
		heap.Node{Data: 12, Priority: 4},
		heap.Node{Data: 15, Priority: 2},
		heap.Node{Data: 17, Priority: 1}}

	h := heap.NewHeap(nos)
	h.Heapfy()
	for _, val := range expectedSequence {
		deleted := h.Delete()
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}

func TestGenericHeapTestIncreasePriorityExample(t *testing.T) {
	nos := []heap.Node{
		heap.Node{Data: 15, Priority: 2},
		heap.Node{Data: 17, Priority: 1},
		heap.Node{Data: 2, Priority: 10},
		heap.Node{Data: 4, Priority: 8},
		heap.Node{Data: 12, Priority: 4},
		heap.Node{Data: 15, Priority: 2},
		heap.Node{Data: 17, Priority: 1},
		heap.Node{Data: 8, Priority: 6},
		heap.Node{Data: 2, Priority: 10}}

	expectedSequence := []heap.Node{
		heap.Node{Data: 2, Priority: 10},
		heap.Node{Data: 17, Priority: 9},
		heap.Node{Data: 4, Priority: 8},
		heap.Node{Data: 8, Priority: 6},
		heap.Node{Data: 12, Priority: 4},
		heap.Node{Data: 15, Priority: 2}}

	h := heap.NewHeap(nos)
	h.Heapfy()
	h.IncreasePriority(heap.Node{Data: 17, Priority: 1}, 9)
	for _, val := range expectedSequence {
		deleted := h.Delete()
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}
