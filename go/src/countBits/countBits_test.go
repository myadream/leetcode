package countBits

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	var n = 2
	var target = []int{0,1,1}
	var result = countBits(n)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("error: %v", result)
	}

	n = 5
	target = []int{0,1,1,2,1,2}
	result = countBits(n)

	if !reflect.DeepEqual(target, result) {
		t.Fatalf("error: %v", result)
	}
}

