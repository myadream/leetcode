package leetcode_2788

import "strings"

func caseOne(words []string, separator byte) (ans []string) {
	ans = []string{}

	for _, word := range words {
		strs := strings.Split(word, string(separator))
		for _, str := range strs {
			if str != "" {
				ans = append(ans, str)
			}
		}

	}

	return ans
}
