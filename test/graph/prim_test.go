package kruskaltest

import (
	"fmt"
	"go_playground/graph"
	"testing"
)

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

	//making undirected graph so adding both edges
	edges := []graph.Edge{
		graph.Edge{S: graph.Node{1}, D: graph.Node{2}, W: 3},
		graph.Edge{S: graph.Node{1}, D: graph.Node{9}, W: 4},
		graph.Edge{S: graph.Node{2}, D: graph.Node{3}, W: 4},
		graph.Edge{S: graph.Node{2}, D: graph.Node{6}, W: 6},
		graph.Edge{S: graph.Node{3}, D: graph.Node{4}, W: 1},
		graph.Edge{S: graph.Node{3}, D: graph.Node{6}, W: 2},
		graph.Edge{S: graph.Node{4}, D: graph.Node{5}, W: 12},
		graph.Edge{S: graph.Node{6}, D: graph.Node{7}, W: 8},
		graph.Edge{S: graph.Node{7}, D: graph.Node{4}, W: 5},
		graph.Edge{S: graph.Node{7}, D: graph.Node{8}, W: 1},
		graph.Edge{S: graph.Node{8}, D: graph.Node{5}, W: 6},
		graph.Edge{S: graph.Node{9}, D: graph.Node{6}, W: 1},
		graph.Edge{S: graph.Node{9}, D: graph.Node{7}, W: 7},
		graph.Edge{S: graph.Node{9}, D: graph.Node{8}, W: 12},
		graph.Edge{S: graph.Node{2}, D: graph.Node{1}, W: 3},
		graph.Edge{S: graph.Node{9}, D: graph.Node{1}, W: 4},
		graph.Edge{S: graph.Node{3}, D: graph.Node{2}, W: 4},
		graph.Edge{S: graph.Node{6}, D: graph.Node{2}, W: 6},
		graph.Edge{S: graph.Node{4}, D: graph.Node{3}, W: 1},
		graph.Edge{S: graph.Node{6}, D: graph.Node{3}, W: 2},
		graph.Edge{S: graph.Node{5}, D: graph.Node{4}, W: 12},
		graph.Edge{S: graph.Node{7}, D: graph.Node{6}, W: 8},
		graph.Edge{S: graph.Node{4}, D: graph.Node{7}, W: 5},
		graph.Edge{S: graph.Node{8}, D: graph.Node{7}, W: 1},
		graph.Edge{S: graph.Node{5}, D: graph.Node{8}, W: 6},
		graph.Edge{S: graph.Node{6}, D: graph.Node{9}, W: 1},
		graph.Edge{S: graph.Node{7}, D: graph.Node{9}, W: 7},
		graph.Edge{S: graph.Node{8}, D: graph.Node{9}, W: 12}}

	g := graph.Graph{}
	g.Nodes = nodes
	g.Edges = edges

	prim := g.Prim(graph.Node{1})
	it := prim.Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
