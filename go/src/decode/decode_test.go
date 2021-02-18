package decode

import (
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	var encoded = []int{1,2,3}
	var first = 1
	var target = []int{1,0,2,1}

	var result = decode(encoded, first)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %d", result)
	}

	encoded = []int{6,2,7,3}
	first = 4
	target = []int{4,2,0,7,4}

	result = decode(encoded, first)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %d", result)
	}

}