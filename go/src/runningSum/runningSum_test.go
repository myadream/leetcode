package runningSum

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	var nums = []int{1,2,3,4}
	var targetResult = []int{1,3,6,10}

	result := runningSum(nums)
	if !reflect.DeepEqual(result, targetResult) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{1,1,1,1,1}
	targetResult = []int{1,2,3,4,5}

	result = runningSum(nums)
	if !reflect.DeepEqual(result, targetResult) {
		t.Fatalf("result: %v", result)
	}

	nums = []int{3,1,2,10,1}
	targetResult = []int{3,4,6,16,17}

	result = runningSum(nums)
	if !reflect.DeepEqual(result, targetResult) {
		t.Fatalf("result: %v", result)
	}


}