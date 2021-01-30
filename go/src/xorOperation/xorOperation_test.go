package xorOperation

import "testing"

func TestXorOperation(t *testing.T) {
	var n = 5
	var start = 0
	var target = 8

	var result = xorOperation(n, start)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	n = 4
	start = 3
	target = 8

	result = xorOperation(n, start)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	n = 1
	start = 7
	target = 7

	result = xorOperation(n, start)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	n = 10
	start = 5
	target = 2

	result = xorOperation(n, start)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}