package numberOfMatches

import "testing"

func TestNumberMatches(t *testing.T) {
	var n = 7
	var target = 6

	var result = numberOfMatches(n)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	n = 14
	target = 13

	result = numberOfMatches(n)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}