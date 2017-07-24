package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestEmptyArray(t *testing.T) {
	if arrayprob.Exists(nil, 1, 2) {
		t.Error("Expected false but returned true")
	}
}

func Test0diffand0distance(t *testing.T) {
	if arrayprob.Exists(nil, 1, 2) {
		t.Error("Expected false but returned true")
	}
}

func TestNegativeDist(t *testing.T) {
	if arrayprob.Exists([]int{1, 2, 3, 4, 5}, 0, -2) {
		t.Error("Expected false but returned true")
	}
}

func TestExistsExample1(t *testing.T) {
	if arrayprob.Exists([]int{1, 2, 3, 4, 5}, 0, -2) {
		t.Error("Expected false but returned true")
	}
}

func TestExistsExample2(t *testing.T) {
	if !arrayprob.Exists([]int{1, 6, 9, 10, 15}, 2, 5) {
		t.Error("Expected true but returned false")
	}
}

func TestExistsExample3(t *testing.T) {
	if !arrayprob.Exists([]int{19, 6, 28, 10, 5}, 2, 6) {
		t.Error("Expected false but returned true")
	}
}

func TestExistsExample4(t *testing.T) {
	if !arrayprob.Exists([]int{8, 6, 1, 17, 15}, 6, 5) {
		t.Error("Expected true but returned false")
	}
}

func TestDoesNotExistsExample4(t *testing.T) {
	if arrayprob.Exists([]int{8, 6, 1, 17, 14}, 2, 5) {
		t.Error("Expected false but returned true")
	}
}
