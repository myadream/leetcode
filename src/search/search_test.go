package search

import "testing"

func TestSearch(t *testing.T) {
	var nums = []int{4,5,6,7,0,1,2}
	var target = 0

	result := search(nums, target)
	if result != target {
		t.Fatalf("result: %d", result)
	}


}