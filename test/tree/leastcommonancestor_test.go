package leastcommonancestortest

import (
	"go_playground/tree"
	"testing"
)

func TestLeastCommonAcestor(t *testing.T) {
	nos := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	tre := tree.CreateFullTree(nos)

	if a, hasLeastCommom := tre.LeastCommonAncestor(8, 9); a != 4 && hasLeastCommom {
		t.Error("expected 4 and true but received ", a, " ", hasLeastCommom)
	}

	if a, hasLeastCommom := tre.LeastCommonAncestor(8, 5); a != 2 && hasLeastCommom {
		t.Error("expected 2 and true but received ", a, " ", hasLeastCommom)
	}

	if a, hasLeastCommom := tre.LeastCommonAncestor(8, 11); a != 2 && hasLeastCommom {
		t.Error("expected 4 and true but received ", a, " ", hasLeastCommom)
	}

	if a, hasLeastCommom := tre.LeastCommonAncestor(8, 2); a != 2 && hasLeastCommom {
		t.Error("expected 4 and true but received ", a, " ", hasLeastCommom)
	}

	// one element on the left of a root and the other on the right
	if a, hasLeastCommom := tre.LeastCommonAncestor(8, 6); a != 1 && hasLeastCommom {
		t.Error("expected 4 and true but received ", a, " ", hasLeastCommom)
	}

	// one element on the left of a root and the other on the right
	if a, hasLeastCommom := tre.LeastCommonAncestor(6, 2); a != 1 && hasLeastCommom {
		t.Error("expected 4 and true but received ", a, " ", hasLeastCommom)
	}

	// one element not found case
	if _, hasLeastCommom := tre.LeastCommonAncestor(18, 7); hasLeastCommom {
		t.Error("expected false but received true ")
	}
}

func TestLeastCommonAcestorEmptyTree(t *testing.T) {
	nos := []int{}
	tre := tree.CreateFullTree(nos)

	if _, hasLeastCommom := tre.LeastCommonAncestor(8, 9); hasLeastCommom {
		t.Error("expected false but received true ")
	}
}

func TestLeastCommonAcestorSingleElementTree(t *testing.T) {
	nos := []int{1}
	tre := tree.CreateFullTree(nos)

	if _, hasLeastCommom := tre.LeastCommonAncestor(1, 9); hasLeastCommom {
		t.Error("expected false but received true ")
	}
}
