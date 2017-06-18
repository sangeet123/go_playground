package graph

// BellmanFord computes single source all node
// shortest path
func (g Graph) BellmanFord(start Node) map[Node]Edge {
	nodeEdgeMap := initializeBellmanFord(g, start)
	for i := 0; i < len(g.Nodes); i++ {
		for _, e := range g.Edges {
			src := e.S
			dest := e.D
			wgt := e.W
			if edge1, ok1 := nodeEdgeMap[src]; ok1 {
				if edge2, ok2 := nodeEdgeMap[dest]; ok2 && edge2.W > edge1.W+wgt {
					nodeEdgeMap[dest] = Edge{S: src, D: dest, W: edge1.W + wgt}
				} else if !ok2 {
					nodeEdgeMap[dest] = Edge{S: src, D: dest, W: edge1.W + wgt}
				}
			}
		}
	}
	return nodeEdgeMap
}

func initializeBellmanFord(g Graph, s Node) map[Node]Edge {
	nodeEdgeMap := make(map[Node]Edge)
	nodeEdgeMap[s] = Edge{W: 0}
	for n, w := range g.getAdjMatrix().neighbors[s] {
		nodeEdgeMap[n] = Edge{S: s, D: n, W: w}
	}
	return nodeEdgeMap
}
