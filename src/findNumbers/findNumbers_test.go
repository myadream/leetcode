package findNumbers

import "testing"

func TestFindNumbers(t *testing.T) {
	var nums = []int{12,345,2,6,7896}
	var target = 2

	var result = findNumbers(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}

	nums = []int{555,901,482,1771}
	target = 1

	result = findNumbers(nums)
	if result != target {
		t.Fatalf("result: %d", result)
	}
}