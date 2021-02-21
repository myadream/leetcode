package isToeplitzMatrix

import "testing"

func TestIsToeplitzMatrix(t *testing.T) {
	var matrix = [][]int{
		{1,2,3,4},
		{5,1,2,3},
		{9,5,1,2},
	}

	var result = isToeplitzMatrix(matrix)
	if result != true {
		t.Fatalf("error")
	}

	matrix = [][]int{
		{1, 2},
		{2, 2},
	}

	result = isToeplitzMatrix(matrix)
	if result != false {
		t.Fatalf("error")
	}
}