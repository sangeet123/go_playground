package multiplyallbutitself


import "testing"
import "reflect"

func TestMultiplyAllButItself(t *testing.T){
	values := []int{1,2,3,4,5,6,7}
	expected := []int{5040,2520,1680,1260,1008,840,720}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received,expected){
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfEmptyArray(t *testing.T){
	values := []int{}
	expected := []int{}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received,expected){
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestMultiplyAllButItselfArrayWithSingleElement(t *testing.T){
	values := []int{12}
	expected := []int{0}

	received := multiply_all_but_itself(values)

	if !reflect.DeepEqual(received,expected){
		t.Error("Expected ", expected, " but received ", received)
	}
}