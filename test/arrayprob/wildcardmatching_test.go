package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestValidWildCardMatchingExample1(t *testing.T) {
	if !arrayprob.MatchesPattern("lcalcalca.c", "lca*lca.c") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample2(t *testing.T) {
	if !arrayprob.MatchesPattern("lcalcalca.c", "lcalca?ca.c") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample3(t *testing.T) {
	if arrayprob.MatchesPattern("abcdegghijjklmn", "ab*g*gghijjklmn") {
		t.Error("Expected false but returned true")
	}
}

func TestValidWildCardMatchingExample4(t *testing.T) {
	if !arrayprob.MatchesPattern("aab", "*ab") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample5(t *testing.T) {
	if !arrayprob.MatchesPattern("abcde", "ab*?de") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample6(t *testing.T) {
	if !arrayprob.MatchesPattern("abcde", "ab****?de") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample7(t *testing.T) {
	if !arrayprob.MatchesPattern("abcde", "ab*******") {
		t.Error("Expected true but returned false")
	}
}

func TestValidWildCardMatchingExample8(t *testing.T) {
	if arrayprob.MatchesPattern("abcde", "ab*ef") {
		t.Error("Expected false but returned true")
	}
}
