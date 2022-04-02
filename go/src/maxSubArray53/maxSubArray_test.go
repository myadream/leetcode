package maxSubArray53

import "testing"

func TestMaxSubArray(t *testing.T) {
	var nums = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	var target = 6

	var result = maxSubArray(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1}
	result = maxSubArray(nums)
	target = 1
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{5, 4, -1, 7, 8}
	target = 23
	result = maxSubArray(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}
