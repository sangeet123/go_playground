package graph

import (
	"go_playground/heap"
	"math"
)

func initializeHeap(g Graph, start Node) *heap.Heap {
	data := []heap.Node{}
	for _, node := range g.Nodes {
		priority := -math.MaxFloat64
		if node.ID == start.ID {
			priority = math.MaxFloat64
		}
		data = append(data, heap.Node{Data: node, Priority: priority})
	}
	h := heap.NewHeap(data)
	h.Heapfy()
	return h
}

// Kruskal returns minimum spanning tree set edges
func (g Graph) Kruskal(start Node) []TraversalResult {
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
		if currentPriority, ok := h.GetPriority(heapNode); ok && currentPriority < -dist {
			h.IncreasePriority(heapNode, dist)
			nodeParentMap[neighbor] = parent
		}
	}
}
