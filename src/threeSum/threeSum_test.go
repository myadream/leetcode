package threeSum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	 nums := []int{-1, 0, 1, 2, -1, -4}
	 target := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	 result := threeSum(nums)

	 if !reflect.DeepEqual(target, result) {
	 	t.Fatalf("result ！= %v, result = %v", target, result)
	 }

	nums = []int{}

	result = threeSum(nums)

	if result != nil {
		t.Fatalf("result ！= nil, result = %v", result)
	}

	nums = []int{0}
	result = threeSum(nums)

	if result != nil {
		t.Fatalf("result ！= nil, result = %v", result)
	}

	nums = []int{1,-1,-1,0}
	target = [][]int{{-1, 0, 1}}

	result = threeSum(nums)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result ！= %v, result = %v", target, result)
	}

	nums = []int{0,0,0,0}
	target = [][]int{{0,0,0}}
	result = threeSum(nums)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result ！= %v, result = %v", target, result)
	}

}