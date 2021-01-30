package combinationSum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var candidates = []int{10,1,2,7,6,1,5}
	var target = [][]int{
		{1, 1, 6},
		{1, 2, 5},
		{1, 7},
		{2, 6},
	}

	var result = combinationSum(candidates, 8)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	candidates = []int{2,5,2,1,2}
	target = [][]int{
		{1, 2, 2},
		{5},
	}

	result = combinationSum(candidates, 5)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}