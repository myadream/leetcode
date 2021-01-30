package pivotIndex

import (
	"fmt"
	"testing"
)

func TestPivotIndex(t *testing.T) {
	var nums = []int{1, 7, 3, 6, 5, 6}
	var target = 3
	var result = pivotIndex(nums)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = []int{1, 2, 3}
	target = -1
	result = pivotIndex(nums)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	fmt.Println(1111)

}
