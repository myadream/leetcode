package NumArray

import "testing"

func TestNumArray_SumRange(t *testing.T) {
	numArray := Constructor([]int{-2, 0, 3, -5, 2, -1})
	result := numArray.SumRange(0, 2)
	if result != 1 {
		t.Fatalf("error")
	}

	result = numArray.SumRange(2, 5)
	if result != -1 {
		t.Fatalf("error")
	}

	result = numArray.SumRange(0, 5)
	if result != -3 {
		t.Fatalf("error")
	}
}
