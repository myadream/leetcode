package flipAndInvertImage

import (
	"reflect"
	"testing"
)

func TestFlipAndInvertImage(t *testing.T) {
	var A = [][]int{
		{1,1,0},
		{1,0,1},
		{0,0,0},
	}

	var target = [][]int{
		{1,0,0},
		{0,1,0},
		{1,1,1},
	}

	var result = flipAndInvertImage(A)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}

	A = [][]int{
		{1,1,0,0},
		{1,0,0,1},
		{0,1,1,1},
		{1,0,1,0},
	}

	 target = [][]int{
		{1,1,0,0},
		{0,1,1,0},
		{0,0,0,1},
		{1,0,1,0},
	}

	 result = flipAndInvertImage(A)
	if !reflect.DeepEqual(target, result) {
		t.Fatalf("result: %v", result)
	}
}