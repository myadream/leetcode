package firstMissingPositive

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	var nums = []int{3, 4, -1, 1, 9, -5}
	var target = 2

	var result = firstMissingPositiveOne(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1,2,0}
	target = 3

	result = firstMissingPositiveOne(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{3,4,-1,1}
	target = 2

	result = firstMissingPositiveOne(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{7,8,9,11,12}
	target = 1

	result = firstMissingPositiveOne(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

}
