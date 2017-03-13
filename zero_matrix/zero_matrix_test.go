package zeromatrix

import "testing"
import "reflect"

func TestZeroMatrixAllZero(t *testing.T){
	input := [][]int{{0,0,0},{1,1,1},{0,0,0}}
	output := [][]int{{0,0,0},{0,0,0},{0,0,0}}
	zero_matrix(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestZeroMatrixFirstColumnThirdRowSet(t *testing.T){
	input := [][]int{{1,0,0},{1,0,0},{1,1,1}}
	output := [][]int{{0,0,0},{0,0,0},{1,0,0}}
	zero_matrix(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestZeroMatrixAllSet(t *testing.T){
	input := [][]int{{1,1,1},{1,1,1},{1,1,1}}
	output := [][]int{{1,1,1},{1,1,1},{1,1,1}}
	zero_matrix(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestZeroMatrixExtremeExampleRowSet(t *testing.T){
	input := [][]int{{1,1,1},{1,1,1},{0,1,1}}
	output := [][]int{{0,1,1},{0,1,1},{0,0,0}}
	zero_matrix(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestZeroMatrixOtherExtremeExampleRowSet(t *testing.T){
	input := [][]int{{0,1,0},{1,1,1},{0,1,0}}
	output := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	zero_matrix(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}


//=======================================Inplace testing================================================================//

func TestInplaceZeroMatrixAllZero(t *testing.T){
	input := [][]int{{0,0,0},{1,1,1},{0,0,0}}
	output := [][]int{{0,0,0},{0,0,0},{0,0,0}}
	zero_matrix_inplace(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestInplaceZeroMatrixFirstColumnThirdRowSet(t *testing.T){
	input := [][]int{{1,0,0},{1,0,0},{1,1,1}}
	output := [][]int{{0,0,0},{0,0,0},{1,0,0}}
	zero_matrix_inplace(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestInplaceZeroMatrixAllSet(t *testing.T){
	input := [][]int{{1,1,1},{1,1,1},{1,1,1}}
	output := [][]int{{1,1,1},{1,1,1},{1,1,1}}
	zero_matrix_inplace(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestInplaceZeroMatrixExtremeExampleRowSet(t *testing.T){
	input := [][]int{{1,1,1},{1,1,1},{0,1,1}}
	output := [][]int{{0,1,1},{0,1,1},{0,0,0}}
	zero_matrix_inplace(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}

func TestInplaceZeroMatrixOtherExtremeExampleRowSet(t *testing.T){
	input := [][]int{{0,1,0},{1,1,1},{0,1,0}}
	output := [][]int{{0,0,0},{0,1,0},{0,0,0}}
	zero_matrix_inplace(input)
	if !reflect.DeepEqual(input, output){
		t.Error("Expected ", output, " but got", input)
	}
}



