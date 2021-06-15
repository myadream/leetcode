package hammingWeight

import "testing"

func TestHammingWeight(t *testing.T) {
	var A uint32 = 00000000000000000000000000001011
	var target = 3
	var result = hammingWeight(A)
	if target != result {
		t.Fatalf("result %d", result)
	}

	A  = 00000000000000000000000010000000
	target = 1
	result = hammingWeight(A)
	if target != result {
		t.Fatalf("result %d", result)
	}

	//A  = 11111111111111111111111111111101
	//target = 31
	//result = hammingWeight(A)
	//if target != result {
	//	t.Fatalf("result %d", result)
	//}

}
