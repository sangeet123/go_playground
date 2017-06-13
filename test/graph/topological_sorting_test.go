package topologialsortingtest

import (
	"fmt"
	"go_playground/graph"
	"testing"
)

func TestTreeInsertOperationWithNoCycles(t *testing.T) {
	nodes := []graph.Node{
		graph.Node{2},
		graph.Node{3},
		graph.Node{5},
		graph.Node{7},
		graph.Node{8},
		graph.Node{9},
		graph.Node{10},
		graph.Node{11}}

	edges := []graph.Edge{
		graph.Edge{S: graph.Node{5}, D: graph.Node{11}},
		graph.Edge{S: graph.Node{11}, D: graph.Node{2}},
		graph.Edge{S: graph.Node{11}, D: graph.Node{9}},
		graph.Edge{S: graph.Node{11}, D: graph.Node{10}},
		graph.Edge{S: graph.Node{7}, D: graph.Node{11}},
		graph.Edge{S: graph.Node{7}, D: graph.Node{8}},
		graph.Edge{S: graph.Node{8}, D: graph.Node{9}},
		graph.Edge{S: graph.Node{3}, D: graph.Node{8}},
		graph.Edge{S: graph.Node{3}, D: graph.Node{10}}}

	g := graph.Graph{}
	g.Nodes = nodes
	g.Edges = edges

	topologicalSorting := g.GetTopologicalOrder()
	fmt.Println(topologicalSorting)
}

func TestTreeInsertOperationWithCycle(t *testing.T) {
	nodes := []graph.Node{
		graph.Node{10},
		graph.Node{2},
		graph.Node{3},
		graph.Node{5},
		graph.Node{10}}

	edges := []graph.Edge{
		graph.Edge{S: graph.Node{2}, D: graph.Node{3}},
		graph.Edge{S: graph.Node{3}, D: graph.Node{5}},
		graph.Edge{S: graph.Node{5}, D: graph.Node{2}},
		graph.Edge{S: graph.Node{1}, D: graph.Node{10}}}

	g := graph.Graph{}
	g.Nodes = nodes
	g.Edges = edges

	defer func() {
		if r := recover(); r == nil {
			t.Error("expected error but no error ocurred")
		}
	}()
	g.GetTopologicalOrder()
}
