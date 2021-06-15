package hammingDistance

import "testing"

func TestHammingDistance(t *testing.T) {
	var x = 1
	var y = 4

	var target = 2
	var result = hammingDistance(x,y)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	x = 93
	y = 73

	target = 2
	result = hammingDistance(x,y)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}