package validcapital


import "testing"

func TestValidCapitalWithAllSmallCase(t *testing.T){
	if !ValidCapital("helloworld"){
		t.Error("Expected true, but got", false)
	}
}

func TestValidCapitalWithAllUpperCase(t *testing.T){
	if !ValidCapital("HELLOWORLD"){
		t.Error("Expected true, but got", false)
	}
}

func TestValidCapitalWithFirstCharUpper(t *testing.T){
	if !ValidCapital("Helloworld"){
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapitalWithWhiteSpace(t *testing.T){
	if ValidCapital("Hello world"){
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapital(t *testing.T){
	if ValidCapital("HelloWorld"){
		t.Error("Expected true, but got", false)
	}
}

func TestInValidCapitalEmptyString(t *testing.T){
	if !ValidCapital(""){
		t.Error("Expected true, but got", false)
	}
}