package url

import "testing"

// Note we have enough space on an array to replace space by %20
// We do not do any validation here
// Validation shoudl be simple enough though
// if there are five spaces, there should be additional 10 extra space at back

func convertToString(url []uint8) string {
	received := ""
	for _, v := range url {
		received = received + string(v)
	}
	return received
}

func TestCharArrayWithSpace(t *testing.T) {
	url := []uint8{'a', ' ', 'b', ' ', 'c', 'e', 'f', 'g', ' ', ' ', ' ', ' '}
	expected := "a%20b%20cefg"
	Urlify(url, len(url))
	received := convertToString(url)
	if expected != received {
		t.Error("Expected ", expected, " but got", received)
	}
}

func TestCharArrayWithoutSpace(t *testing.T) {
	url := []uint8{'a', 'b', 'c', 'e', 'f', 'g'}
	expected := "abcefg"
	Urlify(url, len(url))
	received := convertToString(url)
	if expected != received {
		t.Error("Expected ", expected, " but got", received)
	}
}

func TestEmptyCharArray(t *testing.T) {
	url := []uint8{}
	expected := ""
	Urlify(url, len(url))
	received := convertToString(url)
	if expected != received {
		t.Error("Expected ", expected, " but got", received)
	}
}
