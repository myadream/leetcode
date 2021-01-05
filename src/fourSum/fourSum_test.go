package fourSum

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	nums := []int{1, 0, -1, 0, -2, 2}
	numsTarget := 0
	target := [][]int{
		{-2, -1, 1, 2},
		{-2,  0, 0, 2},
		{-1,  0, 0, 1},
	}

	result := fourSum(nums, numsTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{0, 0, 0, 0}
	numsTarget = 0
	target = [][]int{
		{0, 0, 0, 0},
	}

	result = fourSum(nums, numsTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{-2,-1,-1,1,1,2,2}
	numsTarget = 0
	target = [][]int{
		{-2,-1,1,2},
		{-1,-1,1,1},
	}

	result = fourSum(nums, numsTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}