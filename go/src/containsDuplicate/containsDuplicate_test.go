package containsDuplicate

import "testing"

func TestContainDuplicate(t *testing.T) {
	var A = []int{1, 2, 3, 1}
	var result = containsDuplicate(A)

	if result != true {
		t.Fatalf("error")
	}

	A = []int{1, 2, 3, 4}
	result = containsDuplicate(A)
	if result != false {
		t.Fatalf("error")
	}

	A = []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
	result = containsDuplicate(A)
	if result != true {
		t.Fatalf("error")
	}

	A = []int{0, 4, 5, 0, 3, 6}
	result = containsDuplicate(A)
	if result != true {
		t.Fatalf("error")
	}
}
