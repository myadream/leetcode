package diagonalSum

import "testing"

func TestDiagonalSum(t *testing.T) {
	var nums = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	var target = 25
	var result = diagonalSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = [][]int{
		{1,1,1,1},
		{1,1,1,1},
		{1,1,1,1},
		{1,1,1,1},
	}

	target = 8
	result = diagonalSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	nums = [][]int{
		{5},
	}

	target = 5
	result = diagonalSum(nums)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}
