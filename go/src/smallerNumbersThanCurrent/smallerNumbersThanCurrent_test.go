package smallerNumbersThanCurrent

import (
	"reflect"
	"testing"
)

func TestSmallerNumbersThanCurrentOne(t *testing.T) {
	var nums = []int{8,1,2,2,3}
	var target = []int{4,0,1,1,3}

	var result = smallerNumbersThanCurrentOne(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{6,5,4,8}
	target = []int{2,1,0,3}

	result = smallerNumbersThanCurrentOne(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{7,7,7,7}
	target = []int{0,0,0,0}

	result = smallerNumbersThanCurrentOne(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}

func TestSmallerNumbersThanCurrentTwo(t *testing.T) {
	var nums = []int{8,1,2,2,3}
	var target = []int{4,0,1,1,3}

	var result = smallerNumbersThanCurrentTwo(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{6,5,4,8}
	target = []int{2,1,0,3}

	result = smallerNumbersThanCurrentTwo(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{7,7,7,7}
	target = []int{0,0,0,0}

	result = smallerNumbersThanCurrentTwo(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}