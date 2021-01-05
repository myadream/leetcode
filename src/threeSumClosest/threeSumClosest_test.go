package threeSumClosest

import "testing"

func TestThreeSumClosest(t *testing.T) {
	val := []int{-1,2,1,-4}
	target := 2

	result := threeSumClosest(val, 1)
	if target != result {
		t.Fatalf("error: %d", result)
	}

	val = []int{1,1,1,0}
	target = 3
	result = threeSumClosest(val, 100)
	if target != result {
		t.Fatalf("error: %d", result)
	}
}