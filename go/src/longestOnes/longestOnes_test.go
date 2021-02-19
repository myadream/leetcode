package longestOnes

import "testing"

func TestLongestOnes(t *testing.T) {
	var A = []int{1,1,1,0,0,0,1,1,1,1,0}
	var K = 2
	var target = 6
	var result = longestOnes(A, K)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	A = []int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1}
	K = 3
	target = 10
	result = longestOnes(A, K)

	if target != result {
		t.Fatalf("result: %d", result)
	}

}