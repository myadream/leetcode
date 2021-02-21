package longestSubarray

import "testing"

func TestLongestSubarray(t *testing.T) {
	var num = []int{8,2,4,7}
	var limit = 4
	var target = 2
	var result = longestSubarray(num, limit)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	num = []int{10,1,2,4,7,2}
	limit = 5
	target = 4
	result = longestSubarray(num, limit)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	num = []int{4,2,2,2,4,4,2,2}
	limit = 0
	target = 3
	result = longestSubarray(num, limit)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}
