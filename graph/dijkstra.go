package graph

import "go_playground/heap"

//Dijkstra find the shortest path algorithm of non negative edege cycle graph
// using dijkstra algorithm
func (g Graph) Dijkstra(start Node) map[Node]Node {
	// we consider priority of node which is current
	// distance from a source to update its neighbors priority
	increFunc := func(node heap.Node) float64 {
		return node.Priority
	}
	return g.startProcessing(start, increFunc)
}
