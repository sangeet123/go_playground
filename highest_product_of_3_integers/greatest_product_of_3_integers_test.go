package highestproduceofthreeintegers

import "testing"

func TestHighestProductWithLessThanThreeElements(t *testing.T){
	values := []int{1,2}
	defer func() {
        if r := recover(); r == nil {
            t.Error("no error ocurred")
        }
    }()
	get_highest_product(values)
}

func TestHighestProductWithThanThreeElements(t *testing.T){
	values := []int{1,2,-1}
	expected := -2
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}


func TestHighestProductWithAscendingArray(t *testing.T){
	values := []int{1,2,3,4,5,6,7}
	expected := 210
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}


func TestHighestProductWithDescendingArray(t *testing.T){
	values := []int{7,6,5,3,2,1}
	expected := 210
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}


func TestHighestProductWithAllNegativeArray(t *testing.T){
	values := []int{-1,-10,-15,-2,-3,-20,-1}
	expected := -2
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}


func TestHighestProductWithOnePositiveNegativeArray(t *testing.T){
	values := []int{-1,-10,-15,-2,-3,-20,1}
	expected := 300
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}


func TestHighestProductWithTwoPositiveArray(t *testing.T){
	values := []int{-1,-10,-15,-2,-3,-20,1,2}
	expected := 600
	received := get_highest_product(values)
	if expected != received{
		t.Error("Expected ", expected , " but received ", received)
	}
}