package plusOne

import (
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var digits = []int{1, 2, 3}
	var target = []int{1, 2, 4}
	var result = plusOne(digits)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	digits = []int{4,3,2,1}
	target = []int{4,3,2,2}
	result = plusOne(digits)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	digits = []int{0}
	target = []int{1}
	result = plusOne(digits)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	digits = []int{9}
	target = []int{1,0}
	result = plusOne(digits)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	digits = []int{9,9}
	target = []int{1,0,0}
	result = plusOne(digits)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
	//digits = []int{0,0}
	//target = []int{1}
	//result = plusOne(digits)
	//
	//if !reflect.DeepEqual(target, result) {
	//	t.Fatalf("result: %v", result)
	//}
}