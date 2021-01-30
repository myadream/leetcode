package minimumEffortPath

import "testing"

func TestMinimumEffortPath(t *testing.T) {
	var nums = [][]int{
		{1,2,2},
		{3,8,2},
		{5,3,5},
	}

	var target = 2
	var result = minimumEffortPath(nums)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = [][]int{
		{1,2,3},
		{3,8,4},
		{5,3,5},
	}

	target = 1
	result = minimumEffortPath(nums)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = [][]int{
		{1,2,1,1,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,2,1,2,1},
		{1,1,1,2,1},
	}

	target = 0
	result = minimumEffortPath(nums)

	if target != result {
		t.Fatalf("result: %d", result)
	}
}