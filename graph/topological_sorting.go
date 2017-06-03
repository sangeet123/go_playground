package graph

const (
	visited       = 1
	fullyexplored = 2
)

//GetTopologicalOrder Returns topological sorting of direct acyclic graph
func (g Graph) GetTopologicalOrder() []TraversalResult {
	adjMatrix := g.getAdjMatrix()
	traversedNode := make(map[Node]int)
	sortedOrder := []TraversalResult{}
	for _, node := range g.Nodes {
		if _, ok := traversedNode[node]; !ok {
			traversedNode[node] = visited
			sortedOrder = append([]TraversalResult{TraversalResult{N: node, P: Node{}}}, sortedOrder...)
			topologicalOrder(adjMatrix, node, traversedNode, &sortedOrder)
		}
	}
	return sortedOrder
}

func topologicalOrder(adjMatrix *AdjMatrix, curNode Node, traversedNode map[Node]int, sortedOrder *[]TraversalResult) {
	for neighbor := range adjMatrix.neighbors[curNode] {
		if val, ok := traversedNode[neighbor]; !ok {
			traversedNode[neighbor] = visited
			*sortedOrder = append([]TraversalResult{TraversalResult{N: neighbor, P: curNode}}, *sortedOrder...)
			topologicalOrder(adjMatrix, neighbor, traversedNode, sortedOrder)
		} else if val == visited {
			panic("Not a directed acylic graph")
		}
	}
	traversedNode[curNode] = fullyexplored
}
