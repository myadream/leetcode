package leetcode_20

import "strings"

func CaseOne(s string) bool {

	for i := (len(s) >> 1) + 1; i > 0; i-- {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "[]", "")
		s = strings.ReplaceAll(s, "{}", "")
	}

	return len(s) == 0
}

func CaseTwo(s string) bool {

	if len(s) <= 1 {
		return false
	}

	stack := make([]int32, 0)
	m := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, word := range s {
		if _, ok := m[word]; ok {
			stack = append(stack, m[word])
		} else if len(stack) == 0 || word != stack[len(stack)-1] {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}

	}

	return len(stack) == 0
}
