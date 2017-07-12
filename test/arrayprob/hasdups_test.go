package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestEmptyArray(t *testing.T) {
	arr := []int{}
	if hasDup := arrayprob.HasDups(arr, 1); hasDup {
		t.Error("Expected false but received true")
	}
}

func TestOneElementArray(t *testing.T) {
	arr := []int{1}
	if hasDup := arrayprob.HasDups(arr, 1); hasDup {
		t.Error("Expected false but received true")
	}
}

func TestTwoElementArray(t *testing.T) {
	arr := []int{1, 1}
	if hasDup := arrayprob.HasDups(arr, 1); !hasDup {
		t.Error("Expected true but received false")
	}
}

func TestDupsExample1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 4}
	if hasDup := arrayprob.HasDups(arr, 4); !hasDup {
		t.Error("Expected true but received false")
	}
}

func TestDupsExample2(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 4, 12, 11, 3, 4}
	if hasDup := arrayprob.HasDups(arr, 6); !hasDup {
		t.Error("Expected true but received false")
	}
}

func TestDupsExample3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 4, 7}
	if hasDup := arrayprob.HasDups(arr, 4); hasDup {
		t.Error("Expected false but received true")
	}
}
