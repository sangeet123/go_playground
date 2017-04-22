package teemo

import "testing"

func TestTeemoAttackForContinuousInvervals(t *testing.T) {
	duration := 2
	time_series := []int{1, 3, 6, 8}
	expected := 8
	received := findPoisonedDuration(time_series, duration)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestTeemoAttackForJumpingInvervals(t *testing.T) {
	duration := 2
	time_series := []int{1, 5, 12, 18}
	expected := 8
	received := findPoisonedDuration(time_series, duration)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestTeemoAttackForOverLappingInvervals(t *testing.T) {
	duration := 2
	time_series := []int{1, 2, 3, 4}
	expected := 5
	received := findPoisonedDuration(time_series, duration)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestTeemoAttackForHavingAllContinuousJumpingAndOverLappingIntervals(t *testing.T) {
	duration := 3
	time_series := []int{1, 3, 7, 18, 20}
	expected := 13
	received := findPoisonedDuration(time_series, duration)

	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}
