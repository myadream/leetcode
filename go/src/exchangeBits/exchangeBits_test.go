package exchangeBits

import "testing"

func TestExchangeBits(t *testing.T) {
	var num = 2
	var target = 1
	var result = exchangeBits(num)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	num = 3
	target = 3
	result = exchangeBits(num)
	if target != result {
		t.Fatalf("result: %d", result)
	}


}
