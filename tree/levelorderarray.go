package tree

import (
	"go_playground/list"
)

func (f FullTree) GetLevelOrderArray() [][]int {
	return getLevelOrderArray(f.root)
}

func getLevelOrderArray(n *node) [][]int {
	result := [][]int{}
	list := new(list.List)
	if n == nil {
		return result
	}

	list.Append(n)
	elementsInThisLevel := 1
	for elementsInThisLevel > 0 {
		toAppend := []int{}
		for i := 0; i < elementsInThisLevel; i++ {
			n := list.Get(0).(*node)
			toAppend = append(toAppend, n.data)
			list.Delete(n)
			if n.l != nil {
				list.Append(n.l)
			}
			if n.r != nil {
				list.Append(n.r)
			}
		}
		result = append(result, toAppend)
		elementsInThisLevel = list.Size()
	}
	return result
}
