package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestJavaRegex(t *testing.T) {

	if !arrayprob.MatchesRegex("abc", ".*c") {
		t.Error("Expected true but returned false")
	} //true
	if arrayprob.MatchesRegex("abc", ".*cd") {
		t.Error("Expected false but returned true")
	} // false

	if arrayprob.MatchesRegex("aa", "a") {
		t.Error("Expected false but returned true")
	} //false

	if !arrayprob.MatchesRegex("aa", "aa") {
		t.Error("Expected true but returned false")
	} //true

	if !arrayprob.MatchesRegex("aaa", "a*") {
		t.Error("Expected true but returned false")
	} //true

	if !arrayprob.MatchesRegex("aaaaaaadefg", ".*") {
		t.Error("Expected true but returned false")
	} //true

	if !arrayprob.MatchesRegex("avvvvv==", ".*") {
		t.Error("Expected true but returned false")
	} //true

	if !arrayprob.MatchesRegex("aab", "c*a*b") {
		t.Error("Expected true but returned false")
	} //true

	if arrayprob.MatchesRegex("", ".") {
		t.Error("Expected false but returned true")
	} //false

	if !arrayprob.MatchesRegex("abcdefghd", ".*....") {
		t.Error("Expected true but returned false")
	} //true

	if arrayprob.MatchesRegex("abcdefghd", ".*....dd") {
		t.Error("Expected false but returned true")
	} //false
}
