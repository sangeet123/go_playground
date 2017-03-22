package finddups

import "testing"

func TestFindDuplicatesWithSingleDups(t *testing.T){
	nums := []int{4,3,2,1,1}
	expected := 1
	received := get_first_dups(nums)

	if expected != received{
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestFindDuplicatesWithNoDups(t *testing.T){
	nums := []int{4,3,2,1,5}
	expected := 0
	received := get_first_dups(nums)

	if expected != received{
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestFindDuplicatesWithMultipleDups(t *testing.T){
	nums := []int{4,5,6,10,8,6,5,1,2,3}
	expected := 6
	received := get_first_dups(nums)

	if expected != received{
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestFindDuplicatesWithMultipleDupsAndOneDupsSideBySide(t *testing.T){
	nums := []int{4,5,5,6,10,8,6,5,1,2,3}
	expected := 5
	received := get_first_dups(nums)

	if expected != received{
		t.Error("Expected ", expected, " but received ", received)
	}
}