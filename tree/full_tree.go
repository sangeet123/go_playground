package tree

import (
	"go_playground/list"
)

//This class has been made to create a tree considering
// elements that exist in an array. Each element are added from
// left to right order.
type FullTree struct {
	root *node
}

// CreateFullTree inserts according element on a array
func CreateFullTree(data []int) *FullTree {
	fullTree := getFullTree()
	if len(data) == 0 {
		return fullTree
	}
	list := new(list.List)
	fullTree.root = getNode(data[0])
	list.Append(fullTree.root)

	for i := 1; i < len(data); {
		n := list.Get(0).(*node)
		makeLeftChild(n, getNode(data[i]))
		i++
		if i < len(data) {
			makeRightChild(n, getNode(data[i]))
			i++
		}
		list.Delete(n)
		list.Append(n.l)
		list.Append(n.r)
	}
	return fullTree
}

func getFullTree() *FullTree {
	tree := new(FullTree)
	tree.root = nil
	return tree
}

func (this FullTree) LevelOrder() []int {
	toRet := []int{}
	levelOrder(this.root, &toRet)
	return toRet
}
