package defangIPaddr

import "testing"

func TestDefangIPaddr(t *testing.T) {
	var address = "1.1.1.1"
	var target = "1[.]1[.]1[.]1"
	var result = defangIPaddr(address)

	if target != result {
		t.Fatalf("result: %s", result)
	}

	address = "255.100.50.0"
	target = "255[.]100[.]50[.]0"
	result = defangIPaddr(address)

	if target != result {
		t.Fatalf("result: %s", result)
	}
}
