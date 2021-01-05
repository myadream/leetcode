package searchRange

import (
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	nums := []int{5,7,7,8,8,10}
	needTarget := 8
	target := []int{3, 4}


	result := searchRange(nums, needTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{5,7,7,8,8,10}
	needTarget = 6
	target = []int{-1, -1}


	result = searchRange(nums, needTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{}
	needTarget = 6
	target = []int{-1, -1}

	result = searchRange(nums, needTarget)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}