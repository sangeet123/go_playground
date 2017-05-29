package tree

type intComparator func(a int, b int) int
type intptr *[]int

type node struct {
	data int
	l    *node // left child
	r    *node // right child
	p    *node // parent node
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

// inserts new node with data equals to data
func (this *intTree) Insert(data int) {
	if this.root == nil {
		this.root = getNode(data)
		return
	}
	parent, nodeWithSameData := this.find(data)
	if nodeWithSameData != nil {
		return
	}
	newNode := getNode(data)
	if this.comp(parent.data, data) > 0 {
		makeLeftChild(parent, newNode)
		return
	}
	makeRightChild(parent, newNode)
}

func (this *intTree) find(data int) (*node, *node) {
	start := this.root
	parent := this.root.p
	val := 1
	for start != nil && val != 0 {
		val = this.comp(start.data, data)
		if val == 0 {
			break
		}
		parent = start
		if val < 0 {
			start = start.r
		} else {
			start = start.l
		}
	}
	return parent, start
}

func (this *node) delete() *node {
	var replacer *node
	if this.isLeafNode() {
		createLinkWithGrandParent(this, nil)
		replacer = this.p
	} else if this.hasBothChild() {
		preDecessor := this.getInorderPredecessor()
		createLinkWithGrandParent(preDecessor, preDecessor.r)
		makeLeftChild(preDecessor, this.l)
		makeRightChild(preDecessor, this.r)
		createLinkWithGrandParent(this, preDecessor)
		replacer = preDecessor
	} else if this.hasLeftChild() {
		createLinkWithGrandParent(this, this.l)
		replacer = this.l
	} else {
		createLinkWithGrandParent(this, this.r)
		replacer = this.r
	}
	return replacer
}

func (this *intTree) Delete(data int) {
	_, nodeToDelete := this.find(data)
	if nodeToDelete == nil {
		return
	}
	replacer := nodeToDelete.delete()
	if nodeToDelete.isRootNode() {
		this.root = replacer
		this.root.p = nil
	}
	nodeToDelete.stripAllLinks()
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

func (this intTree) GetHeight() int {
	return height(this.root)
}

func height(n *node) int {
	if n == nil {
		return 0
	}
	leftHeight := 1 + height(n.l)
	rightHeight := 1 + height(n.r)

	if leftHeight >= rightHeight {
		return leftHeight
	}
	return rightHeight
}
