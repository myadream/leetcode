package majorityElement

import "testing"

func TestMajorityElement(t *testing.T) {
	target := []int{1,2,5,9,5,9,5,5,5}

	result := majorityElement(target)
	if result != 5 {
		t.Fatal("result not eq 5")
	}

	target = []int{3, 2}
	result = majorityElement(target)
	if result != -1 {
		t.Fatal("result not eq -1")
	}

	target = []int{2,2,1,1,1,2,2}
	result = majorityElement(target)
	if result != 2 {
		t.Fatal("result not eq 2")
	}
}
