package createTargetArray

import (
	"reflect"
	"testing"
)

func TestCreateTargetArray(t *testing.T) {
	var nums = []int{0, 1, 2, 3, 4}
	var index = []int{0, 1, 2, 2, 1}
	var target = []int{0, 4, 1, 3, 2}

	var result = createTargetArray(nums, index)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{1, 2, 3, 4, 0}
	index = []int{0, 1, 2, 3, 0}
	target = []int{0,1,2,3,4}

	result = createTargetArray(nums, index)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{1}
	index = []int{0}
	target = []int{1}

	result = createTargetArray(nums, index)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
}