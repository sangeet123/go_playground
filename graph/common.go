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
	if nodes[i].Priority <= nodes[j].Priority {
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

func updateHeapAndParentMap(h *heap.Heap, parent heap.Node, neighbors map[Node]float64, nodeParentMap map[Node]Edge, incFunc increaseBy) {
	for neighbor, dist := range neighbors {
		heapNode := heap.Node{Data: neighbor}
		if curDist, ok := h.GetPriority(heapNode); ok {
			if priority, ok := updatePriority(curDist, dist, incFunc(parent)); ok {
				h.IncreasePriority(heapNode, priority)
				nodeParentMap[neighbor] = Edge{S: parent.Data.(Node), D: neighbor, W: priority}
			}
		}
	}
}

func (g Graph) startProcessing(start Node, incFunc increaseBy) map[Node]Edge {
	h := initializeHeap(g, start)
	nodeParentMap := make(map[Node]Edge)
	nodeParentMap[start] = Edge{S: start, D: Node{}, W: 0}
	for !h.IsEmpty() {
		node := h.Delete()
		neighbors := g.getAdjMatrix().neighbors[node.Data.(Node)]
		updateHeapAndParentMap(h, node, neighbors, nodeParentMap, incFunc)
	}
	return nodeParentMap
}

func updatePriority(cur float64, to float64, inc float64) (float64, bool) {
	if cur > inc+to {
		return inc + to, true
	}
	return cur, false
}
