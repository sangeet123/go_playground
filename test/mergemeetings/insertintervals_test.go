package mergemeetingstest

import (
	"go_playground/mergemeetings"
	"reflect"
	"testing"
)

func TestNonOverlappingIntervalsExample1(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 8}, {Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 0, End: 1}, {Start: 2, End: 3}, {Start: 4, End: 8}, {Start: 10, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 0, End: 1})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample2(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4,End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 6, End: 11})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample3(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 30}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 21, End: 30})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample4(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 15}, {Start: 19, End: 21}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 2, End: 14})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample5(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 13}, {Start: 14, End: 15}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 10, End: 25}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 10, End: 25})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample6(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 12}, {Start: 14, End: 18}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 18}, {Start: 19, End: 21}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 12, End: 14})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample7(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 12}, {Start: 14, End: 18}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 12}, {Start: 14, End: 30}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 18, End: 30})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestNonOverlappingIntervalsExample8(t *testing.T) {
	ranges := mergemeetings.Meetings{{Start: 2, End: 3}, {Start: 4, End: 6}, {Start: 11, End: 12}, {Start: 14, End: 18}, {Start: 19, End: 21}}
	expected := mergemeetings.Meetings{{Start: 1, End: 21}}
	received := mergemeetings.InsertInterval(ranges, mergemeetings.Meeting{Start: 1, End: 21})

	if !reflect.DeepEqual(expected, received) {
		t.Error("Expected ", expected, " but received ", received)
	}
}
