package removeOuterParentheses

import "testing"

func TestRemoveOuterParentheses(t *testing.T) {
	var S = "(()())(())"
	var target = "()()()"

	var result = removeOuterParentheses(S)
	if target != result {
		t.Fatalf("result: ----%s----", result)
	}

	S = "(()())(())(()(()))"
	target = "()()()()(())"

	result = removeOuterParentheses(S)
	if target != result {
		t.Fatalf("result: %s", result)
	}

	S = "()()"
	target = ""

	result = removeOuterParentheses(S)
	if target != result {
		t.Fatalf("result: %s", result)
	}
}