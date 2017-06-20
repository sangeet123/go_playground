package graph

import (
	"math"
)

// FloydWarshall computes all pair shortest
// path algorithm
func (g Graph) FloydWarshall() (map[Node]map[Node]float64, map[Node]map[Node]Node) {
	distMatrix, parentMatrix := g.initializeFloydWarshall()
	for _, n1 := range g.Nodes {
		for _, n2 := range g.Nodes {
			for _, n3 := range g.Nodes {
				distn1n2 := distMatrix[n1][n2]
				distn3n2 := distMatrix[n3][n2]
				distn1n3 := distMatrix[n1][n3]
				if updateDist(distn1n2, distn1n3, distn3n2) {
					distMatrix[n1][n2] = distn1n3 + distn3n2
					parentMatrix[n1][n2] = n3
				}
			}
		}
	}
	return distMatrix, parentMatrix
}

func updateDist(x, y, z float64) bool {
	if y == math.MaxFloat64 || z == math.MaxFloat64 {
		return false
	}

	if x == math.MaxFloat64 {
		return true
	}

	if x > y+z {
		return true
	}

	return false
}

func (g Graph) initializeFloydWarshall() (map[Node]map[Node]float64, map[Node]map[Node]Node) {
	distMatrix := make(map[Node]map[Node]float64)
	parentMatrix := make(map[Node]map[Node]Node)
	adjMatrix := g.getAdjMatrix()
	// initialize distance matrix
	for _, n1 := range g.Nodes {
		for _, n2 := range g.Nodes {
			dist, parent := getDistance(n1, n2, adjMatrix)
			if _, ok := distMatrix[n1]; !ok {
				distMatrix[n1] = make(map[Node]float64)
				parentMatrix[n1] = make(map[Node]Node)
			}
			distMatrix[n1][n2] = dist
			parentMatrix[n1][n2] = parent
		}
	}
	return distMatrix, parentMatrix
}

func getDistance(n1, n2 Node, adjMatrix *AdjMatrix) (float64, Node) {
	if n1 == n2 {
		return 0.0, Node{}
	} else if neighbor, ok := adjMatrix.neighbors[n1]; ok {
		if _, ok := neighbor[n2]; ok {
			return neighbor[n2], n1
		}
	}
	return math.MaxFloat64, Node{}
}
