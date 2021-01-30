package combinationSum

import (
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	var candidates = []int{2,3,6,7}
	var target = [][]int{
		{2, 2, 3},
		{7},
	}

	var result = combinationSum(candidates, 7)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}

	candidates = []int{2,3,5}
	target = [][]int{
		{2,2,2,2},
		{2,3,3},
		{3,5},
	}

	result = combinationSum(candidates, 8)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v", result)
	}
}