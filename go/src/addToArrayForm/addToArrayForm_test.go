package addToArrayForm

import (
	"reflect"
	"testing"
)

func TestAddToArrayForm(t *testing.T) {
	var A = []int{1,2,0,0}
	var K = 34

	var target = []int{1,2,3,4}

	var result = addToArrayForm(A, K)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v",result)
	}

	A = []int{2,7,4}
	K = 181

	target = []int{4,5,5}

	result = addToArrayForm(A, K)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v",result)
	}

	A = []int{2,1,5}
	K = 806

	target = []int{1,0,2,1}

	result = addToArrayForm(A, K)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v",result)
	}

	A = []int{9,9,9,9,9,9,9,9,9,9}
	K = 1

	target = []int{1,0,0,0,0,0,0,0,0,0,0}

	result = addToArrayForm(A, K)
	if !reflect.DeepEqual(result, target) {
		t.Fatalf("result: %v",result)
	}
}