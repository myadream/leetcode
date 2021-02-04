package characterReplacement

import "testing"

func TestCharacterReplacement(t *testing.T) {
	var s = "ABAB"
	var k = 2
	var target = 4

	var result = characterReplacement(s, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	s = "AABABBA"
	k = 1
	target = 4

	result = characterReplacement(s, k)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}