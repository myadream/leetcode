package equalSubstring

import "testing"

func TestEqualSubstring(t *testing.T) {
	var s = "abcd"
	var st = "bcdf"
	var consts = 3
	var target = 3

	var result = equalSubstring(s, st, consts)
	if target != result {
		t.Fatalf("result: %d",result)
	}

	s = "abcd"
	st = "cdef"
	consts = 3
	target = 1

	result = equalSubstring(s, st, consts)
	if target != result {
		t.Fatalf("result: %d",result)
	}


	s = "abcd"
	st = "acde"
	consts = 0
	target = 1

	result = equalSubstring(s, st, consts)
	if target != result {
		t.Fatalf("result: %d",result)
	}

	s = "krrgw"
	st = "zjxss"
	consts = 19
	target = 2

	result = equalSubstring(s, st, consts)
	if target != result {
		t.Fatalf("result: %d",result)
	}
}