package fairCandySwap

import (
	"reflect"
	"testing"
)

func TestFairCandySwap(t *testing.T) {
	var A = []int{1, 1}
	var B = []int{2, 2}

	var target = []int{1, 2}
	var result = fairCandySwap(A, B)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	A = []int{1, 1}
	B = []int{2, 3}

	target = []int{1, 2}
	result = fairCandySwap(A, B)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	A = []int{2}
	B = []int{1, 3}

	target = []int{2, 3}
	result = fairCandySwap(A, B)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	A = []int{1, 2, 5}
	B = []int{2, 4}

	target = []int{5, 4}
	result = fairCandySwap(A, B)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
}