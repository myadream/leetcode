package checkPossibility

import "testing"

func TestCheckPossibility(t *testing.T) {
	var nums = []int{4,2,3}
	var result = checkPossibility(nums)

	if true != result {
		t.Fatalf("result: %v", result)
	}

	nums = []int{4,2,1}
	result = checkPossibility(nums)

	if false != result {
		t.Fatalf("result: %v", result)
	}
}