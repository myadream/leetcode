package minTimeToVisitAllPoints

import (
	"testing"
)

func TestMinTimeToVisitAllPoints(t *testing.T) {
	var nums = [][]int{
		{1, 1},
		{3, 4},
		{-1, 0},
	}

	var target = 7
	var result = minTimeToVisitAllPoints(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	nums = [][]int{
		{3, 2},
		{-2, 2},
	}

	target = 5
	result = minTimeToVisitAllPoints(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}
