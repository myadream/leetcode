package searchInsert

import "testing"

func TestSearchInsert(t *testing.T) {
	nums := []int{1,3,5,6}
	target := 2

	result := searchInsert(nums, 5)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	target = 1
	result = searchInsert(nums, 2)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	target = 4
	result = searchInsert(nums, 7)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	target = 0
	result = searchInsert(nums, 0)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}
