package decompressRLElist

import (
	"reflect"
	"testing"
)

func TestDecompressRLElist(t *testing.T) {
	var nums = []int{1,2,3,4}
	var target = []int{2,4,4,4}

	var result = decompressRLElist(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{1,1,2,3}
	target = []int{1,3,3}

	result = decompressRLElist(nums)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}