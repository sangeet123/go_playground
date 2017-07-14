package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestTriangleExample1(t *testing.T) {
	arr := [][]int{}
	if sum := arrayprob.GetMin(arr); sum != 0 {
		t.Error("Expected 0 but received ", sum)
	}
}

func TestTriangleExample2(t *testing.T) {
	arr := [][]int{{1}, {2, 1}, {3, 4, 1}, {4, 3, 2, 1}, {5, 4, 3, 2, 1}}
	if sum := arrayprob.GetMin(arr); sum != 5 {
		t.Error("Expected 5 but received ", sum)
	}
}

func TestTriangleExample3(t *testing.T) {
	arr := [][]int{{1}, {2, 1}, {3, 4, 1}, {4, 3, 2, 1}, {4, 4, 3, 2, 3}}
	if sum := arrayprob.GetMin(arr); sum != 6 {
		t.Error("Expected 5 but received ", sum)
	}
}

func TestTriangleExample4(t *testing.T) {
	arr := [][]int{{5}, {1, 8}, {10, 12, 3}, {8, 6, 2, 4}, {2, 3, 1, 4, 5}}
	if sum := arrayprob.GetMin(arr); sum != 19 {
		t.Error("Expected 5 but received ", sum)
	}
}
