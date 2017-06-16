package graph

import (
	"fmt"
	"go_playground/list"
	"sort"
)

type unionset struct {
	n     Node
	count int
}

// Len calculates length of edges
func (edges Edges) Len() int {
	return len(edges)
}

// Less checkes which of the two edges has less distance
func (edges Edges) Less(i, j int) bool {
	return edges[i].W < edges[j].W
}

// Swap swaps the edges
func (edges Edges) Swap(i, j int) {
	edges[i], edges[j] = edges[j], edges[i]
}

func initialize(nodes []Node) map[Node]unionset {
	setMap := make(map[Node]unionset)
	// Initialization
	for _, n := range nodes {
		newSetItem := unionset{n: n}
		setMap[n] = newSetItem
	}
	return setMap
}

// Kruskal algorithm for finding minimum spanning tree
func (g Graph) Kruskal() *list.List {
	kruskalEdges := new(list.List)
	setMap := initialize(g.Nodes)
	sort.Sort(g.Edges)
	for _, e := range g.Edges {
		val := belongsToSameSet(setMap, e.S, e.D)
		fmt.Println("Edge:", e, " val :=", val, " setMap", setMap)
		if !val {
			merge(setMap, e.S, e.D)
			kruskalEdges.Append(e)
		}
	}
	return kruskalEdges
}

func merge(setMap map[Node]unionset, s, d Node) {
	sSet := getSet(setMap, s)
	dSet := getSet(setMap, d)

	if sSet.count >= dSet.count {
		sSet.count = sSet.count + dSet.count + 1
		setMap[dSet.n] = sSet
		return
	}
	dSet.count = dSet.count + sSet.count + 1
	setMap[sSet.n] = dSet
}

func belongsToSameSet(setMap map[Node]unionset, s, d Node) bool {
	return getSet(setMap, s) == getSet(setMap, d)
}

func getSet(setMap map[Node]unionset, n Node) unionset {
	if setMap[n].n == n {
		return setMap[n]
	}
	s := getSet(setMap, setMap[n].n)
	setMap[n] = s
	return s
}
