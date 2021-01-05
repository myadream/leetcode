package removeDuplicates

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1,1,2}
	targetLen := 2

	result := removeDuplicates(nums)
	if result != targetLen {
		t.Fatalf("result: %d", result)
	}

	nums = []int{0,0,1,1,1,2,2,3,3,4}
	targetLen = 5
	result = removeDuplicates(nums)
	if result != targetLen {
		t.Fatalf("result: %d", result)
	}
}