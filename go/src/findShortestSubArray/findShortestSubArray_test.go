package findShortestSubArray

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	var A = []int{1, 2, 2, 3, 1}
	var target = 2
	var result = findShortestSubArray(A)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	A = []int{1,2,2,3,1,4,2}
	target = 6
	result = findShortestSubArray(A)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}