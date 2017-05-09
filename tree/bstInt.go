package tree

type intComparator func(a int, b int) int
type intptr *[]int

type node struct {
	data int
	l    *node
	r    *node
}

type intTree struct {
	root *node
	comp intComparator
}

func GetIntTree(c intComparator) *intTree {
	tree := new(intTree)
	tree.root = nil
	tree.comp = c
	return tree
}

func getNode(data int) *node {
	n := new(node)
	n.data = data
	n.l = nil
	n.r = nil
	return n
}

func (this *intTree) Insert(data int) {
	newNode := getNode(data)
	if this.root == nil {
		this.root = newNode
		return
	}

	start := this.root
	parent := start

	for start != nil {
		parent = start
		if this.comp(start.data, data) >= 0 {
			start = start.l
		} else {
			start = start.r
		}
	}

	if this.comp(parent.data, data) > 0 {
		parent.l = newNode
	} else {
		parent.r = newNode
	}
}

func (this intTree) InOrder() []int {
	toRet := []int{}
	inOrder(this.root, &toRet)
	return toRet
}

func inOrder(n *node, toRetPtr intptr) {
	if n != nil {
		inOrder(n.l, toRetPtr)
		*toRetPtr = append(*toRetPtr, n.data)
		inOrder(n.r, toRetPtr)
	}
}

func (this intTree) PreOrder() []int {
	toRet := []int{}
	preOrder(this.root, &toRet)
	return toRet
}

func preOrder(n *node, toRetPtr intptr) {
	if n != nil {
		*toRetPtr = append(*toRetPtr, n.data)
		preOrder(n.l, toRetPtr)
		preOrder(n.r, toRetPtr)
	}
}

func (this intTree) PostOrder() []int {
	toRet := []int{}
	postOrder(this.root, &toRet)
	return toRet
}

func postOrder(n *node, toRetPtr intptr) {
	if n != nil {
		postOrder(n.l, toRetPtr)
		postOrder(n.r, toRetPtr)
		*toRetPtr = append(*toRetPtr, n.data)
	}
}
