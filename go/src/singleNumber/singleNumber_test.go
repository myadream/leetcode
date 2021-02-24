package singleNumber

import "testing"

func TestSingleNuber(t *testing.T) {
	var N = []int{2,2,1}
	var target = 1
	var result = singleNumber(N)

	if target != result {
		t.Fatalf("reslt: %d", result)
	}

	N = []int{4,1,2,1,2}
	target = 4
	result = singleNumber(N)

	if target != result {
		t.Fatalf("reslt: %d", result)
	}
}