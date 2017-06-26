package graph

import (
	"go_playground/list"
  //"fmt"
)

// HasRoute method checks if there is a route
// between a and b. This is unidirectional search
func(g Graph) HasRoute(a Node, b Node) bool{
  adjMatrix := g.getAdjMatrix()
  traversedNode := make(map[Node]struct{})
  traversedNode[a] = struct{}{}
  list := new(list.List)
  list.Append(a)

  for list.Size() > 0 {
		node0 := list.Get(0).(Node)
		list.Delete(node0)
		for neighbor := range adjMatrix.neighbors[node0] {
    if neighbor == b{
        return true
      }
			if _, ok := traversedNode[neighbor]; !ok {
				traversedNode[neighbor] = struct{}{}
				list.Append(neighbor)
			}
		}
	}
  return false
}

//Bidirectional search
func(g Graph) FastHasRoute(a Node, b Node) bool{
  var adjMatrices [2]*AdjMatrix
  adjMatrices[0]=g.getAdjMatrix()
  adjMatrices[1]=g.getTransposeAdjMatrix()


  traversedNode := [2]map[Node]struct{}{make(map[Node]struct{}),make(map[Node]struct{})}
  traversedNode[0][a] = struct{}{}
  traversedNode[1][b] = struct{}{}

  var lists [2]*list.List
  lists[0] = new(list.List)
  lists[1] = new(list.List)
  lists[0].Append(a)
  lists[1].Append(b)

  i := 0
  for lists[0].Size() > 0 || lists[1].Size() > 0{
    if lists[i].Size() <= 0{
      i = (i + 1)%2
      continue
    }
    node0 := lists[i].Get(0).(Node)
    neighbors := adjMatrices[i].neighbors[node0]
		lists[i].Delete(node0)
		for neighbor := range neighbors {
      j := (i + 1) % 2

      if _, ok := traversedNode[j][neighbor]; ok {
        return true
      }

      if _, ok := traversedNode[i][neighbor]; !ok {
				traversedNode[i][neighbor] = struct{}{}
				lists[i].Append(neighbor)
			}
		}
    i = (i + 1) % 2
	}
  return false
}
