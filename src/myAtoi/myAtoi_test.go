package myAtoi

import "testing"


func TestMyAtoi(t *testing.T) {
	var input = " 42"
	var target = 42

	var result = myAtoi(input)
	if target != result {
		t.Fatalf("result: %v", result)
	}


	var input = " 42"
	var target = 42

	var result = myAtoi(input)
	if target != result {
		t.Fatalf("result: %v", result)
	}

}