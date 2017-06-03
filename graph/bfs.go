package graph

import (
	"go_playground/list"
)

// GetBFSTraversal does bredth first traversal
func (g Graph) GetBFSTraversal(start Node) []TraversalResult {
	adjMatrix := g.getAdjMatrix()
	traversedNode := make(map[Node]struct{})

	// Initialization
	bfsResultSet := []TraversalResult{TraversalResult{N: start, P: Node{}}}
	traversedNode[start] = struct{}{}
	list := new(list.List)
	list.Append(start)

	for list.Size() > 0 {
		node0 := list.Get(0).(Node)
		list.Delete(node0)
		for neighbor := range adjMatrix.neighbors[node0] {
			if _, ok := traversedNode[neighbor]; !ok {
				traversedNode[neighbor] = struct{}{}
				bfsResultSet = append(bfsResultSet, TraversalResult{N: neighbor, P: node0})
				list.Append(neighbor)
			}
		}
	}
	return bfsResultSet
}
