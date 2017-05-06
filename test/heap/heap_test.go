package heaptest

import (
	"go_playground/heap"
	"reflect"
	"testing"
)

func TestHeapTest(t *testing.T) {
	nos := []int{5, 4, 3, 2, 1}
	expected := []int{1, 2, 3, 5, 4}
	heap.HeapfyInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}

	expected = []int{1, 2, 3, 4, 5}
	for _, val := range expected {
		deleted := heap.Delete(&nos)
		if deleted != val {
			t.Error("Expected ", val, " but received ", deleted)
		}
	}
}

func TestHeapTestSingleElementArray(t *testing.T) {
	nos := []int{1}
	expected := []int{1}
	heap.HeapfyInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestHeapTestTwoElementsArray(t *testing.T) {
	nos := []int{2, 1}
	expected := []int{1, 2}
	heap.HeapfyInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}

func TestHeapTestExample(t *testing.T) {
	nos := []int{15, 17, 2, 4, 12, 15, 17, 8, 2}
	expected := []int{2, 4, 2, 8, 12, 15, 17, 15, 17}
	heap.HeapfyInt(nos)
	if !reflect.DeepEqual(expected, nos) {
		t.Error("Expected ", expected, " but received ", nos)
	}
}
