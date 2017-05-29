package dfstest

import (
	"fmt"
	"go_playground/graph"
	"testing"
)

type node graph.Node
type edge graph.Edge

func TestTreeInsertOperation(t *testing.T) {
	nodes := []graph.Node{
		graph.Node{1},
		graph.Node{2},
		graph.Node{3},
		graph.Node{4},
		graph.Node{5},
		graph.Node{6},
		graph.Node{7},
		graph.Node{8},
		graph.Node{9}}

	edges := []graph.Edge{
		graph.Edge{S: graph.Node{1}, D: graph.Node{2}},
		graph.Edge{S: graph.Node{1}, D: graph.Node{9}},
		graph.Edge{S: graph.Node{2}, D: graph.Node{3}},
		graph.Edge{S: graph.Node{2}, D: graph.Node{6}},
		graph.Edge{S: graph.Node{3}, D: graph.Node{4}},
		graph.Edge{S: graph.Node{3}, D: graph.Node{6}},
		graph.Edge{S: graph.Node{4}, D: graph.Node{5}},
		graph.Edge{S: graph.Node{6}, D: graph.Node{7}},
		graph.Edge{S: graph.Node{7}, D: graph.Node{4}},
		graph.Edge{S: graph.Node{7}, D: graph.Node{8}},
		graph.Edge{S: graph.Node{8}, D: graph.Node{5}},
		graph.Edge{S: graph.Node{9}, D: graph.Node{6}},
		graph.Edge{S: graph.Node{9}, D: graph.Node{7}},
		graph.Edge{S: graph.Node{9}, D: graph.Node{8}}}

	g := graph.Graph{}
	g.Nodes = nodes
	g.Edges = edges

	bfsTraversal := g.GetBFSTraversal(graph.Node{1})
	fmt.Println(bfsTraversal)
}
