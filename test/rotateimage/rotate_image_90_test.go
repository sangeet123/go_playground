package rotateimagetest

import (
	"go_playground/rotateimage"
	"reflect"
	"testing"
)

func TestRotateImageEmptyArray(t *testing.T) {
	input := [][]int{}
	output := [][]int{}
	received := rotateimage.RotateImgBy90(input)

	if len(received) != 0 {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageSingleElementArray(t *testing.T) {
	input := [][]int{{1}}
	output := [][]int{{1}}
	received := rotateimage.RotateImgBy90(input)

	if !reflect.DeepEqual(output, received) {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageSingleRowArray(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}}
	output := [][]int{{1}, {2}, {3}, {4}}
	received := rotateimage.RotateImgBy90(input)

	if !reflect.DeepEqual(output, received) {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageSingleColumnArray(t *testing.T) {
	input := [][]int{{1}, {2}, {3}, {4}}
	output := [][]int{{4, 3, 2, 1}}
	received := rotateimage.RotateImgBy90(input)

	if !reflect.DeepEqual(output, received) {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageMultiDimensionalArray(t *testing.T) {
	input := [][]int{{1, 2, 3}, {4, 5, 6}}
	output := [][]int{{4, 1}, {5, 2}, {6, 3}}
	received := rotateimage.RotateImgBy90(input)

	if !reflect.DeepEqual(output, received) {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageMultiDimensionalSquareArrayEvenDimension(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	output := [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}
	received := rotateimage.RotateImgBy90(input)

	if !reflect.DeepEqual(output, received) {
		t.Error("Expected ", output, " but got ", received)
	}
}

func TestRotateImageInplaceMultiOneElementSquareArray(t *testing.T) {
	input := [][]int{{1}}
	output := [][]int{{1}}
	rotateimage.RotateImgBy90InPlace(input)
	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageInplaceMultiDimensionalSquareArrayOddDimension(t *testing.T) {
	//input
	//1,2,3                7,4,1
	//4,5,6    =========>  8,5,2
	//7,8,9				   9,6,3
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	rotateimage.RotateImgBy90InPlace(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageInplaceMultiDimensionalSquareArrayEvenDimension(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	output := [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}
	rotateimage.RotateImgBy90InPlace(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageInplaceMultiDimensionalSquareArrayOddSevenDimension(t *testing.T) {
	// input
	// 1,  2,   3,  4,  5,  6,  7       			 15, 8,  1, 22,  15,  8,  1
	// 8,  9,   10, 11, 12, 13, 14    	=========>   16, 9,  2, 23,  16,  9,  2
	// 15, 16,  17, 18, 19, 20, 21				     17, 10, 3, 24,  17,  10, 3
	// 22, 23,  24, 25, 26, 27, 28					 18, 11, 4, 25,  18,  11  4
	// 1,  2,   3,  4,  5,  6,  7       			 19, 12, 5, 26,  19,  12, 5
	// 8,  9,   10, 11, 12, 13, 14    	=========>   20, 13, 6, 27,  20,  13, 6
	// 15, 16,  17, 18, 19, 20, 21				     21, 14, 7, 28,  21,  14, 7

	input := [][]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}, {22, 23, 24, 25, 26, 27, 28}, {1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}}
	output := [][]int{{15, 8, 1, 22, 15, 8, 1}, {16, 9, 2, 23, 16, 9, 2}, {17, 10, 3, 24, 17, 10, 3}, {18, 11, 4, 25, 18, 11, 4}, {19, 12, 5, 26, 19, 12, 5}, {20, 13, 6, 27, 20, 13, 6}, {21, 14, 7, 28, 21, 14, 7}}
	rotateimage.RotateImgBy90InPlace(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

//============================================O(2n) less efficient ===============================================================

func TestRotateImageUsingTransposeMultiOneElementSquareArray(t *testing.T) {
	input := [][]int{{1}}
	output := [][]int{{1}}
	rotateimage.RotateImgBy90InPlaceUsingTranspose(input)
	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageUsingTransposeMultiDimensionalSquareArrayOddDimension(t *testing.T) {
	//input
	//1,2,3                7,4,1
	//4,5,6    =========>  8,5,2
	//7,8,9				   9,6,3
	input := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	output := [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}
	rotateimage.RotateImgBy90InPlaceUsingTranspose(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageUsingTransposeMultiDimensionalSquareArrayEvenDimension(t *testing.T) {
	input := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	output := [][]int{{13, 9, 5, 1}, {14, 10, 6, 2}, {15, 11, 7, 3}, {16, 12, 8, 4}}
	rotateimage.RotateImgBy90InPlaceUsingTranspose(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateImageUsingTransposeMultiDimensionalSquareArrayOddSevenDimension(t *testing.T) {
	// input
	// 1,  2,   3,  4,  5,  6,  7       			 15, 8,  1, 22,  15,  8,  1
	// 8,  9,   10, 11, 12, 13, 14    	=========>   16, 9,  2, 23,  16,  9,  2
	// 15, 16,  17, 18, 19, 20, 21				     17, 10, 3, 24,  17,  10, 3
	// 22, 23,  24, 25, 26, 27, 28					 18, 11, 4, 25,  18,  11  4
	// 1,  2,   3,  4,  5,  6,  7       			 19, 12, 5, 26,  19,  12, 5
	// 8,  9,   10, 11, 12, 13, 14    	=========>   20, 13, 6, 27,  20,  13, 6
	// 15, 16,  17, 18, 19, 20, 21				     21, 14, 7, 28,  21,  14, 7

	input := [][]int{{1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}, {22, 23, 24, 25, 26, 27, 28}, {1, 2, 3, 4, 5, 6, 7}, {8, 9, 10, 11, 12, 13, 14}, {15, 16, 17, 18, 19, 20, 21}}
	output := [][]int{{15, 8, 1, 22, 15, 8, 1}, {16, 9, 2, 23, 16, 9, 2}, {17, 10, 3, 24, 17, 10, 3}, {18, 11, 4, 25, 18, 11, 4}, {19, 12, 5, 26, 19, 12, 5}, {20, 13, 6, 27, 20, 13, 6}, {21, 14, 7, 28, 21, 14, 7}}
	rotateimage.RotateImgBy90InPlaceUsingTranspose(input)

	if !reflect.DeepEqual(output, input) {
		t.Error("Expected ", output, " but got ", input)
	}
}

func TestRotateInplaceImageSingleRowArray(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected panic but failed")
		}
	}()
	input := [][]int{{1, 2, 3, 4}}
	rotateimage.RotateImgBy90InPlaceUsingTranspose(input)
}
