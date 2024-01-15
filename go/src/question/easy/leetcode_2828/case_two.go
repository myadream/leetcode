package leetcode_2828

func CaseTwo(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	str := ""
	for _, word := range words {
		str += string(word[0])
	}

	return str == s && str != ""
}
