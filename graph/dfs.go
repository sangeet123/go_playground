package graph

// GetDFSTraversal performs dfs traversal of graph
func (g Graph) GetDFSTraversal() []TraversalResult {
	adjMatrix := g.getAdjMatrix()
	traversedNode := make(map[Node]struct{})
	dfsResultSet := []TraversalResult{}
	for _, node := range g.Nodes {
		if _, ok := traversedNode[node]; !ok {
			traversedNode[node] = struct{}{}
			dfsResultSet = append(dfsResultSet, TraversalResult{N: node, P: Node{}})
			dfsTraversal(adjMatrix, node, traversedNode, &dfsResultSet)
		}
	}
	return dfsResultSet
}

func dfsTraversal(adjMatrix *AdjMatrix, curNode Node, traversedNode map[Node]struct{}, dfsResultSet *[]TraversalResult) {
	for neighbor := range adjMatrix.neighbors[curNode] {
		if _, ok := traversedNode[neighbor]; !ok {
			traversedNode[neighbor] = struct{}{}
			*dfsResultSet = append(*dfsResultSet, TraversalResult{N: neighbor, P: curNode})
			dfsTraversal(adjMatrix, neighbor, traversedNode, dfsResultSet)
		}
	}
}
