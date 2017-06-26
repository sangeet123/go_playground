package medianofsortedarrtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestArrMidElements(t *testing.T) {
	nos := []int{4, 3, 2, 1, 10, 11, 12, 13, 5, 6, 7, 8, 9}
	if val := arrayprob.GetKthLargestElement(nos, 6); val != 6 {
		t.Error("Expected 6 but received ", val)
	}
}

func TestFirstElements(t *testing.T) {
	nos := []int{4, 3, 2, 1, 10, 11, 12, 13, 5, 6, 7, 8, 9}
	if val := arrayprob.GetKthLargestElement(nos, 1); val != 1 {
		t.Error("Expected 1 but received ", val)
	}
}

func TestThirteenElements(t *testing.T) {
	nos := []int{4, 3, 2, 1, 10, 11, 12, 13, 5, 6, 7, 8, 9}
	if val := arrayprob.GetKthLargestElement(nos, 13); val != 13 {
		t.Error("Expected 6 but received ", val)
	}
}
