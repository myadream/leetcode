package numberOfSteps

import "testing"

func TestNumberOfSteps (t *testing.T)  {
	var num = 14
	var target = 6
	var result = numberOfSteps(num)
	if target != result {
		t.Fatalf("result: %d", result)
	}

	num = 8
	target = 4
	result = numberOfSteps(num)
	if target != result {
		t.Fatalf("result: %d", result)
	}
}