package graph

import (
  "fmt"
  "sort"
)

type nodeInfo struct {
  n Node
	f int // finish time
}

type nodesInfo []nodeInfo

//for sorting implement Len(), Less(i,j int) and Swap(i,j int)
func (slice nodesInfo) Len() int {
	return len(slice)
}

func (slice nodesInfo) Less(i, j int) bool {
	return slice[i].f > slice[j].f
}

func (slice nodesInfo) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

func (g Graph) GetStronglyConnectedComponent() {
	infoNodes := preProcess(g)
  sort.Sort(infoNodes)
  fmt.Println(infoNodes)
}

func preProcess(g Graph) nodesInfo{
	adjMatrix := g.getAdjMatrix()
	traversedNode := make(map[Node]int)
	startTime := 1
	for _, node := range g.Nodes {
		if _, ok := traversedNode[node]; !ok {
			traversedNode[node] = startTime
			recordStartAndFinishTimeForDFSTraversal(adjMatrix, node, traversedNode, &startTime)
		}
	}
	return populateNodeInfo(traversedNode)
}

func populateNodeInfo(traversedNode map[Node]int)[]nodeInfo{
  result := nodesInfo{}
  for k,v := range traversedNode{
    result = append(result,nodeInfo{k,v})
  }
  return result
}

func recordStartAndFinishTimeForDFSTraversal(adjMatrix *AdjMatrix, curNode Node, traversedNode map[Node]int, time *int) {
	for neighbor := range adjMatrix.neighbors[curNode] {
		if _, ok := traversedNode[neighbor]; !ok {
      *time++
      traversedNode[neighbor] = *time
			recordStartAndFinishTimeForDFSTraversal(adjMatrix, neighbor, traversedNode, time)
		}
	}
  *time++
	traversedNode[curNode] = *time
}
