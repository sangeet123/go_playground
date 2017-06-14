package graph

import (
	"go_playground/heap"
	"math"
)

type increaseBy func(node heap.Node) float64

func compare(nodes []heap.Node, i int, j int) int {
	if i < 0 || i >= len(nodes) || j < 0 || j >= len(nodes) {
		panic("index out of bound")
	}
	if nodes[i].Priority < nodes[j].Priority {
		return i
	}
	return j
}

func initializeHeap(g Graph, start Node) *heap.Heap {
	data := []heap.Node{}
	for _, node := range g.Nodes {
		priority := math.MaxFloat64
		if node.ID == start.ID {
			priority = 0
		}
		data = append(data, heap.Node{Data: node, Priority: priority})
	}
	h := heap.NewHeap(data, compare)
	h.Heapfy()
	return h
}

func updateHeapAndParentMap(h *heap.Heap, parent Node, neighbors map[Node]float64, nodeParentMap map[Node]Node, incFunc increaseBy) {
	for neighbor, dist := range neighbors {
		heapNode := heap.Node{Data: neighbor}
		if curDist, ok := h.GetPriority(heapNode); ok {
			if priority, ok := updatePriority(curDist, dist, incFunc(heapNode)); ok {
				h.IncreasePriority(heapNode, priority)
				nodeParentMap[neighbor] = parent
			}
		}
	}
}

func (g Graph) startProcessing(start Node, incFunc increaseBy) map[Node]Node {
	h := initializeHeap(g, start)
	nodeParentMap := make(map[Node]Node)
	nodeParentMap[start] = Node{}
	for !h.IsEmpty() {
		node := h.Delete()
		nextNode := node.Data.(Node)
		neighbors := g.getAdjMatrix().neighbors[nextNode]
		updateHeapAndParentMap(h, nextNode, neighbors, nodeParentMap, incFunc)
	}
	return nodeParentMap
}

func updatePriority(cur float64, to float64, inc float64) (float64, bool) {
	if cur > inc+to {
		return inc + to, true
	}
	return cur, false
}
