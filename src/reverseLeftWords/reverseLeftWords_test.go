package reverseLeftWords

import "testing"

func TestReverseLeftWords(t *testing.T)  {
	var s = "abcdefg"
	var k = 2
	var target = "cdefgab"

	var result = reverseLeftWords(s, k)
	if target != result {
		t.Fatalf("result: %s", result)
	}

	s = "lrloseumgh"
	k = 6
	target = "umghlrlose"

	result = reverseLeftWords(s, k)
	if target != result {
		t.Fatalf("result: %s", result)
	}
}