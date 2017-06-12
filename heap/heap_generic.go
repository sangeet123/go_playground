package heap

import "fmt"

// CustomData is for storing any data
type CustomData interface{}

//DataTable store heap data
type DataTable []CustomData

// Node of a heap
type Node struct {
	Data     CustomData
	Priority int
}

// Heap data struct
type Heap struct {
	Data []Node
	hmap map[string]int
}

// NewHeap returns the heap data structure
// Here map only stores the unique one
// all data with same val removed
func NewHeap(data []Node) *Heap {
	if len(data) == 0 {
		panic("Data cannot be nil")
	}
	h := new(Heap)
	h.Data = []Node{}
	h.hmap = make(map[string]int)
	i := 0
	for _, val := range data {
		key := getKey(val)
		if _, ok := h.hmap[key]; !ok {
			h.hmap[key] = i
			h.Data = append(h.Data, val)
			i++
		}
	}
	return h
}

// Heapfy that performs integer heapification
func (h Heap) Heapfy() {
	size := len(h.Data)
	for i := size >> 1; i >= 0; i-- {
		h.PerformHeapify(i)
	}
}

// PerformHeapify performs heapification
func (h Heap) PerformHeapify(from int) {
	size := len(h.Data)
	loop := true
	for j, l, r := updateIndex(from); l < size && loop; {
		ind := l

		if r < size {
			ind = h.comp(r, l)
		}

		if h.comp(ind, j) == ind {
			h.Data[ind], h.Data[j] = h.Data[j], h.Data[ind]
			h.hmap[getKey(h.Data[ind])], h.hmap[getKey(h.Data[j])] = h.hmap[getKey(h.Data[j])], h.hmap[getKey(h.Data[ind])]
			j, l, r = updateIndex(ind)
		} else {
			loop = false
		}
	}
}

// Delete removes element from top of heap
func (h *Heap) Delete() Node {
	if len(h.Data) == 0 {
		panic("Empty heap")
	}

	toReturn := h.Data[0]
	h.Data[0] = h.Data[len(h.Data)-1]
	h.Data = h.Data[:len(h.Data)-1]
	delete(h.hmap, getKey(toReturn))
	h.PerformHeapify(0)
	return toReturn
}

// IncreasePriority increases priority for that node
func (h *Heap) IncreasePriority(node Node, priority int) {
	if _, ok := h.hmap[getKey(node)]; ok {
		nodeIndx := h.hmap[getKey(node)]
		h.Data[nodeIndx].Priority = priority
		parent := nodeIndx >> 1
		for ; parent != nodeIndx && h.Data[parent].Priority < h.Data[nodeIndx].Priority; parent = parent >> 1 {
			h.Data[parent], h.Data[nodeIndx] = h.Data[nodeIndx], h.Data[parent]
			h.hmap[getKey(h.Data[parent])], h.hmap[getKey(h.Data[nodeIndx])] = h.hmap[getKey(h.Data[nodeIndx])], h.hmap[getKey(h.Data[parent])]
			nodeIndx = parent
		}
	}
}

func (h Heap) comp(i, j int) int {
	if h.Data[i].Priority >= h.Data[j].Priority {
		return i
	}
	return j
}

func getKey(node Node) string {
	return fmt.Sprintf("%v", node.Data)
}
