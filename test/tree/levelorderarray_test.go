package levelorderarrtest

import (
	"go_playground/tree"
	"reflect"
	"testing"
)

func TestLevelOrderArray(t *testing.T) {
	nos := []int{4, 5, 3, 2, 7, 8, 1, 2}
	result := [][]int{{4}, {5, 3}, {2, 7, 8, 1}, {2}}
	tre := tree.CreateFullTree(nos)

	levelOrder := tre.GetLevelOrderArray()
	if !reflect.DeepEqual(result, levelOrder) {
		t.Error("expected levelorder ", result, " but received ", levelOrder)
	}
}
