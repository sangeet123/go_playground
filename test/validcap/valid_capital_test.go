package validcaptest

import (
	"go_playground/validcap"
	"testing"
)

func TestValidCapitalWithAllSmallCase(t *testing.T) {
	if !validcap.ValidCapital("helloworld") {
		t.Error("Expected true, but got", false)
	}
}

func TestValidCapitalWithAllUpperCase(t *testing.T) {
	if !validcap.ValidCapital("HELLOWORLD") {
		t.Error("Expected true, but got", false)
	}
}

func TestValidCapitalWithFirstCharUpper(t *testing.T) {
	if !validcap.ValidCapital("Helloworld") {
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapitalWithWhiteSpace(t *testing.T) {
	if validcap.ValidCapital("Hello world") {
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapital(t *testing.T) {
	if validcap.ValidCapital("HelloWorld") {
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapitalEmptyString(t *testing.T) {
	if !validcap.ValidCapital("") {
		t.Error("Expected true, but got", false)
	}
}
