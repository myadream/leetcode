package leetcode_2864

import "strings"

func caseOne(s string) string {
	count := 0
	for _, word := range s {
		if word == '1' {
			count++
		}
	}

	n := len(s)
	s = ""

	for i := 0; i < n-1; i++ {
		if count > 1 {
			s += "1"
			count -= 1
		} else {
			s += "0"
		}
	}

	if count == 1 {
		s += "1"
	}

	return s
}

func caseTwo(s string) string {
	count := strings.Count(s, "1")
	return strings.Repeat("1", count-1) + strings.Repeat("0", len(s)-count) + "1"
}
