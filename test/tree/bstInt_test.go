package bstinttest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func comparator(a, b int) int {
	if a == b {
		return 0
	} else if a > b {
		return 1
	}
	return -1
}

func TestTreeInsertOperation(t *testing.T) {
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	expectedInOrder := []int{1, 2, 3, 4, 5, 7, 8}
	expectedPreOrder := []int{4, 3, 2, 1, 5, 7, 8}
	expectedPostOrder := []int{1, 2, 3, 8, 7, 5, 4}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}

}

// Testing for root nodes
func TestTreeDeleteRootWithRightChildOnly(t *testing.T) {
	nos := []int{55, 47}
	expectedInOrder := []int{47}
	expectedPreOrder := []int{47}
	expectedPostOrder := []int{47}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(55)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteRootWithLeftChildOnly(t *testing.T) {
	nos := []int{55, 57}
	expectedInOrder := []int{57}
	expectedPreOrder := []int{57}
	expectedPostOrder := []int{57}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(55)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}

}

func TestTreeDeleteRootWithBothChildAndImmediateInOrderPredecessor(t *testing.T) {
	nos := []int{55, 47, 57, 60}
	expectedInOrder := []int{47, 57, 60}
	expectedPreOrder := []int{57, 47, 60}
	expectedPostOrder := []int{47, 60, 57}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(55)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}

}

func TestTreeDeleteRootWithBothChildAndNonImmediateInOrderPredecessor(t *testing.T) {
	nos := []int{55, 47, 58, 60, 56, 57}
	expectedInOrder := []int{47, 56, 57, 58, 60}
	expectedPreOrder := []int{56, 47, 58, 57, 60}
	expectedPostOrder := []int{47, 57, 60, 58, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(55)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

// Testing for non root nodes
func TestTreeDeleteLeftChildNonRootWithRightChildOnly(t *testing.T) {
	nos := []int{56, 47, 50}
	expectedInOrder := []int{50, 56}
	expectedPreOrder := []int{56, 50}
	expectedPostOrder := []int{50, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(47)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteLeftChildNonRootWithLeftChildOnly(t *testing.T) {
	nos := []int{56, 47, 46}
	expectedInOrder := []int{46, 56}
	expectedPreOrder := []int{56, 46}
	expectedPostOrder := []int{46, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(47)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteRightChildNonRootWithRightChildOnly(t *testing.T) {
	nos := []int{56, 67, 68}
	expectedInOrder := []int{56, 68}
	expectedPreOrder := []int{56, 68}
	expectedPostOrder := []int{68, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(67)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteRightChildNonRootWithLeftChildOnly(t *testing.T) {
	nos := []int{56, 67, 60}
	expectedInOrder := []int{56, 60}
	expectedPreOrder := []int{56, 60}
	expectedPostOrder := []int{60, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(67)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteRightChildNonRootWithImmediatePredecessor(t *testing.T) {
	nos := []int{56, 47, 68, 65, 70, 75}
	expectedInOrder := []int{47, 56, 65, 70, 75}
	expectedPreOrder := []int{56, 47, 70, 65, 75}
	expectedPostOrder := []int{47, 65, 75, 70, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(68)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteLeftChildNonRootWithImmediatePredecessor(t *testing.T) {
	nos := []int{56, 47, 68, 45, 50, 52}
	expectedInOrder := []int{45, 50, 52, 56, 68}
	expectedPreOrder := []int{56, 50, 45, 52, 68}
	expectedPostOrder := []int{45, 52, 50, 68, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(47)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteRightChildNonRootWithNonImmediatePredecessor(t *testing.T) {
	nos := []int{56, 47, 68, 65, 70, 75, 69}
	expectedInOrder := []int{47, 56, 65, 69, 70, 75}
	expectedPreOrder := []int{56, 47, 69, 65, 70, 75}
	expectedPostOrder := []int{47, 65, 75, 70, 69, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(68)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}

func TestTreeDeleteLeftChildNonRootWithNonImmediatePredecessor(t *testing.T) {
	nos := []int{56, 47, 68, 45, 50, 48, 52}
	expectedInOrder := []int{45, 48, 50, 52, 56, 68}
	expectedPreOrder := []int{56, 48, 45, 50, 52, 68}
	expectedPostOrder := []int{45, 52, 50, 48, 68, 56}
	tre := tree.GetIntTree(comparator)
	for _, val := range nos {
		tre.Insert(val)
	}

	tre.Delete(47)

	inorder := tre.InOrder()
	if !reflect.DeepEqual(expectedInOrder, inorder) {
		t.Error("expected inOrder ", expectedInOrder, " but received ", inorder)
	}

	preOrder := tre.PreOrder()
	if !reflect.DeepEqual(expectedPreOrder, preOrder) {
		t.Error("expected preOrder ", expectedPreOrder, " but received ", preOrder)
	}

	postOrder := tre.PostOrder()
	if !reflect.DeepEqual(expectedPostOrder, postOrder) {
		t.Error("expected postOrder ", expectedPostOrder, " but received ", postOrder)
	}
}
