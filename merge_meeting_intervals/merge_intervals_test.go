package mergeMeetings

import "testing"
import "reflect"

func TestNonOverlappingNonSortedIntervals(t *testing.T){
	ranges := Meetings{{start:10,end:13},{start:14,end:15},{start:4,end:8},{start:1,end:3},{start:19, end:21}}
	expected := Meetings{{start:1,end:3},{start:4,end:8},{start:10,end:13},{start:14,end:15},{start:19, end:21}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestNonOverlappingSortedIntervals(t *testing.T){
	ranges := Meetings{{start:1,end:3},{start:4,end:8},{start:10,end:13},{start:14,end:15},{start:19, end:21}}
	expected := Meetings{{start:1,end:3},{start:4,end:8},{start:10,end:13},{start:14,end:15},{start:19, end:21}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestOverlappingSortedIntervals(t *testing.T){
	ranges := Meetings{{start:1,end:4},{start:4,end:8},{start:7,end:13},{start:14,end:15},{start:14, end:21}}
	expected := Meetings{{start:1,end:13},{start:14,end:21}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestOverlappingNonSortedIntervals(t *testing.T){
	ranges := Meetings{{start:14, end:21},{start:4,end:8},{start:1,end:4},{start:14,end:15},{start:7,end:13}}
	expected := Meetings{{start:1,end:13},{start:14,end:21}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestExampleFromCakeInterviewExample1(t *testing.T){
	ranges := Meetings{{start:1, end:2},{start:2,end:3}}
	expected := Meetings{{start:1,end:3}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestExampleFromCakeInterviewExample2(t *testing.T){
	ranges := Meetings{{start:1, end:5},{start:1,end:3}}
	expected := Meetings{{start:1,end:5}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestExampleFromCakeInterviewExample3(t *testing.T){
	ranges := Meetings{{start:1, end:5},{start:2,end:3}}
	expected := Meetings{{start:1,end:5}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}

func TestExampleFromCakeInterviewExample4(t *testing.T){
	ranges := Meetings{{start:1, end:10},{start:2,end:6},{start:3,end:5},{start:7,end:9}}
	expected := Meetings{{start:1,end:10}}
	received := MergeMeetings(ranges)

	if !reflect.DeepEqual(expected,received){
		t.Error("Expected ", expected, " but received ", received);
	}
}