package findComplement

import "testing"

func TestFindComplement(t *testing.T) {
	var n = 5
	var target = 2
	var result = findComplement(n)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}