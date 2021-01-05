package shuffle

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	var nums = []int{2,5,1, 3,4,7}
	var n = 3
	var target = []int{2,3,5,4,1,7}

	var result = shuffle(nums, n)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{1,2,3,4, 4,3,2,1}
	n = 4
	target = []int{1,4,2,3,3,2,4,1}

	result = shuffle(nums, n)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}


	nums = []int{1,1,2,2}
	n = 2
	target = []int{1,2,1,2}

	result = shuffle(nums, n)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
}