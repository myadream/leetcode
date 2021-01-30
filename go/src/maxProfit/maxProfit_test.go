package maxProfit

import "testing"

func TestMaxProfit(t *testing.T) {
	var nums = []int{7,1,5,3,6,4}
	var target = 7

	var result = maxProfit(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1,2,3,4,5}
	target = 4

	result = maxProfit(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	nums = []int{7,6,4,3,1}
	target = 0

	result = maxProfit(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}