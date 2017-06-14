package graph

import "go_playground/heap"

// Prim algorithm only works for directed graph
// and may not find the spanning tree for the directed graph

// Here is an example
//      b
//   /\    \
// 3/       \ 1   5 + 3 = 8 (will be the solution starting from verted s)
// /        \/    5 + 1 = 6 ( is the minimum spanning tree)
// s ------->a
//      5

// Prim returns minimum spanning tree set edges
func (g Graph) Prim(start Node) map[Node]Node {
	// we do not consider priority of node to calculate
	// the priority of neighbors
	increFunc := func(node heap.Node) float64 {
		return 0.0
	}
	return g.startProcessing(start, increFunc)
}
