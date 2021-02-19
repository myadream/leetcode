package intersection

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	var nums1 = []int{1,2,2,1}
	var nums2 = []int{2,2}
	var target = []int{2}

	var result = intersection(nums1, nums2)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result %v", result)
	}

	nums1 = []int{4,9,5}
	nums2 = []int{9,4,9,8,4}
	target = []int{9, 4}

	result = intersection(nums1, nums2)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result %v", result)
	}

}