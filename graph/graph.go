package graph

// Node of a graph
// For the sake of simplicity each node of graph is given unique id
type Node struct {
	Id int
}

// Edge of a graph
type Edge struct {
	S Node    // represents source
	D Node    // represents destination
	W float64 // represents connection weight between two nodes
}

// Graph
type Graph struct {
	Nodes []Node
	Edges []Edge
}

//Graph Representation
type AdjMatrix struct {
	neighbors map[Node]map[Node]float64
}

//Result of traversal
type TraversalResult struct {
	N Node // Represents Node itself
	P Node // Represents Parent Node
}

// graph representation is only needed for internal purpose so not exposing
func (g Graph) getAdjMatrix() *AdjMatrix {
	adjMatrix := new(AdjMatrix)
	adjMatrix.neighbors = make(map[Node]map[Node]float64)
	for _, edge := range g.Edges {
		if _, ok := adjMatrix.neighbors[edge.S]; !ok {
			adjMatrix.neighbors[edge.S] = make(map[Node]float64)
		}
		adjMatrix.neighbors[edge.S][edge.D] = edge.W
	}
	return adjMatrix
}
