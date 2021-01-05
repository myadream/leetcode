package removeElement

import "testing"

//https://leetcode-cn.com/problems/remove-element/
func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	target := 2

	result := removeElement(nums, val)
	if target != result {
		t.Fatalf("result: %v", result)
	}

	nums = []int{0, 1, 2, 2, 3, 0, 4, 2}
	val = 2
	target = 5

	result = removeElement(nums, val)
	if target != result {
		t.Fatalf("result: %v", result)
	}
}