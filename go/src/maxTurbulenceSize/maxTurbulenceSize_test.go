package maxTurbulenceSize

import "testing"

func TestMaxTurbulenceSize(t *testing.T) {
	var A = []int{9,4,2,10,7,8,8,1,9}
	var target = 5
	var result = maxTurbulenceSize(A)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	A = []int{4,8,12,16}
	target = 2
	result = maxTurbulenceSize(A)

	if target != result {
		t.Fatalf("result: %d", result)
	}

	A = []int{100}
	target = 1
	result = maxTurbulenceSize(A)

	if target != result {
		t.Fatalf("result: %d", result)
	}
}
