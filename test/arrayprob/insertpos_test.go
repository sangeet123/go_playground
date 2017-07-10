package arrayprob

import (
	"go_playground/arrayprob"
	"testing"
)

func TestFourSumExample1(t *testing.T) {
	arr := []int{12, 18, 19, 26, 30}
	toInsert := []int{17, 25, 26, 29, 35, 1}
	expectedPositions := []int{1, 3, 3, 4, 5, 0}

	for i := 0; i < len(toInsert); i++ {
		if pos := arrayprob.GetInsertionPos(arr, toInsert[i]); pos != expectedPositions[i] {
			t.Error("Expected ", expectedPositions[i], " for number ", toInsert[i], " but received ", pos)
		}
	}
}
