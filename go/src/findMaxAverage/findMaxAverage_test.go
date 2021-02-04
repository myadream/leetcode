package findMaxAverage

import "testing"

func TestFindMaxAverage(t *testing.T) {
	var nums = []int{1,12,-5,-6,50,3}
	var k = 4
	var target = 12.75

	var result = findMaxAverage(nums, k)
	if target != result {
		t.Fatalf("result: %f", result)
	}
}