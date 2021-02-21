package arrayPairSum

import "testing"

func TestArrayPairSum(t *testing.T) {
	var nums = []int{1,4,3,2}
	var target = 4

	var result = arrayPairSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{6,2,6,5,1,2}
	target = 9

	result = arrayPairSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{7,3,1,0,0,6}
	target = 7

	result = arrayPairSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}