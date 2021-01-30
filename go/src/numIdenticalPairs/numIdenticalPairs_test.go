package numIdenticalPairs

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	var nums = []int{1, 2, 3, 1, 1, 3}
	var target = 4

	var result = numIdenticalPairs(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1, 1, 1, 1}
	target = 6

	result = numIdenticalPairs(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1, 2, 3}
	target = 0

	result = numIdenticalPairs(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}
