package template

func CaseOne(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}

	for i := len(s) - 1; i >= 0; i-- {
		if words[i][0] == 0 || s[i] == 0 || words[i][0] != s[i] {
			return false
		}
	}

	return true
}
