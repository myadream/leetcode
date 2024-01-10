package leetcode_2696

import "strings"

func caseOne(s string) int {
	for i := (len(s) / 2) + 1; i > 0; i-- {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}

	return len(s)
}

func caseTwo(s string) int {
	var stack []byte

	for i := 0; i < len(s); i++ {
		l := len(stack) - 1
		if l >= 0 && (s[i] == 'B' && stack[l] == 'A' || s[i] == 'D' && stack[l] == 'C') {
			stack = stack[:l]
		} else {
			stack = append(stack, s[i])
		}
	}

	return len(stack)
}
