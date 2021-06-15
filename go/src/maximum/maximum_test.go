package maximum

import "testing"

func TestMaximum(t *testing.T) {
	var a = 1
	var b = 2
	var target = 2
	var result = maximum(a, b)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	a = 2147483647
	b = -2147483648
	target = 2147483647
	result = maximum(a, b)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}
