package graph

import (
	"go_playground/heap"
	"math"
)

// Prim algorithm only works for directed graph
// and may not find the spanning tree for the directed graph

// Here is an example
//      b
//   /\    \
// 3/       \ 1   5 + 3 = 8 (will be the solution starting from verted s)
// /        \/    5 + 1 = 6 ( is the minimum spanning tree)
// s ------->a
//      5
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

// Prim returns minimum spanning tree set edges
func (g Graph) Prim(start Node) []TraversalResult {
	h := initializeHeap(g, start)
	nodeParentMap := make(map[Node]Node)
	minimumSpanningResultSet := []TraversalResult{}
	nodeParentMap[start] = Node{}
	for !h.IsEmpty() {
		node := h.Delete()
		nextNode := node.Data.(Node)
		minimumSpanningResultSet = append(minimumSpanningResultSet, TraversalResult{N: nextNode, P: nodeParentMap[nextNode]})
		neighbors := g.getAdjMatrix().neighbors[nextNode]
		updateHeapAndParentMap(h, nextNode, neighbors, nodeParentMap)
	}
	return minimumSpanningResultSet
}

func updateHeapAndParentMap(h *heap.Heap, parent Node, neighbors map[Node]float64, nodeParentMap map[Node]Node) {
	for neighbor, dist := range neighbors {
		heapNode := heap.Node{Data: neighbor}
		if curDist, ok := h.GetPriority(heapNode); ok && curDist > dist {
			h.IncreasePriority(heapNode, dist)
			nodeParentMap[neighbor] = parent
		}
	}
}
