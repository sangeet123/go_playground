package mergemeetingstest

import (
	"go_playground/mergemeetings"
	"reflect"
	"testing"
)

func TestNonOverlappingNonSortedIntervals(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 4, End: 8}, {Start: 1, End: 3}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 1, End: 3}, {Start: 4, End: 8}, {Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingSortedIntervals(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 3}, {Start: 4, End: 8}, {Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 1, End: 3}, {Start: 4, End: 8}, {Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestOverlappingSortedIntervals(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 4}, {Start: 4, End: 8}, {Start: 7, End: 13}, {Start: 14, End: 15}, {Start: 14, End: 21}}
	expected := mergemeetings.Meetings{{Start: 1, End: 13}, {Start: 14, End: 21}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestOverlappingNonSortedIntervals(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 14, End: 21}, {Start: 4, End: 8}, {Start: 1, End: 4}, {Start: 14, End: 15}, {Start: 7, End: 13}}
	expected := mergemeetings.Meetings{{Start: 1, End: 13}, {Start: 14, End: 21}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestExampleFromCakeInterviewExample1(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 2}, {Start: 2, End: 3}}
	expected := mergemeetings.Meetings{{Start: 1, End: 3}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestExampleFromCakeInterviewExample2(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 5}, {Start: 1, End: 3}}
	expected := mergemeetings.Meetings{{Start: 1, End: 5}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestExampleFromCakeInterviewExample3(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 5}, {Start: 2, End: 3}}
	expected := mergemeetings.Meetings{{Start: 1, End: 5}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestExampleFromCakeInterviewExample4(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 1, End: 10}, {Start: 2, End: 6}, {Start: 3, End: 5}, {Start: 7, End: 9}}
	expected := mergemeetings.Meetings{{Start: 1, End: 10}}
	received := mergemeetings.MergeMeetings(ranges)

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}
