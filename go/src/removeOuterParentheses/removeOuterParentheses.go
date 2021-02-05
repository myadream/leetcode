package removeOuterParentheses

import "bytes"

//https://leetcode-cn.com/problems/remove-outermost-parentheses/
func removeOuterParentheses(S string) string {
	var sByte bytes.Buffer
	var sum int

	for k, v := range S {
		if v == 41 {
			sum--
		}

		if sum >= 1	 {
			sByte.WriteByte(S[k])
		}

		if v == 40 {
			sum++
		}
	}

	return sByte.String()
}
