package arrayprobtest

import (
	"go_playground/arrayprob"
	"testing"
)

func TestLongestValidParenExample1(t *testing.T) {
	expected := "((({})))"
	received := arrayprob.LongestValidParen("()(((({})))")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestLongestValidParenExample2(t *testing.T) {
	expected := "(())()"
	received := arrayprob.LongestValidParen(")))(())(){")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestLongestValidParenExample3(t *testing.T) {
	expected := "[][][][][][][][][]()()"
	received := arrayprob.LongestValidParen(")))(())(){}][][][][][][][][][]()()}}[]]}")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestLongestValidParenExample4(t *testing.T) {
	expected := "(()())"
	received := arrayprob.LongestValidParen("(()())")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}

func TestLongestValidParenExample5(t *testing.T) {
	expected := "()"
	received := arrayprob.LongestValidParen("()((()]")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}


func TestLongestValidParenExample6(t *testing.T) {
  expected := "([][]())"
  received := arrayprob.LongestValidParen("()()()]([][]())")
  if expected != received {
    t.Error("Expected ", expected, " but received ", received)
  }
}

func TestLongestValidParenExample7(t *testing.T) {
	expected := ""
	received := arrayprob.LongestValidParen("]]]")
	if expected != received {
		t.Error("Expected ", expected, " but received ", received)
	}
}
