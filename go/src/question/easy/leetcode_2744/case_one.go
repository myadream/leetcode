package leetcode_2744

func caseOne(words []string) (ans int) {
	var hashMap = map[string]uint8{}

	for i := len(words) - 1; i >= 0; i-- {
		hashMap[words[i]] += 1
	}

	for word, num := range hashMap {
		revKey := string(word[1]) + string(word[0])
		nextWord := hashMap[revKey]
		if nextWord == 0 || num == 0 || (word[0] == word[1] && num <= 1) {
			continue
		}

		if word[0] == word[1] {
			ans += int(num >> 1)
		} else {
			ans += min(num, nextWord)
		}

		hashMap[word] = 0
		hashMap[revKey] = 0
	}

	return ans
}

func min(a, b uint8) int {
	if a >= b {
		return int(b)
	}

	return int(a)
}

func caseTwo(words []string) (ans int) {
	seen := map[int]int{}
	for _, word := range words {
		ans += seen[int(word[1])*100+int(word[0])]
		seen[int(word[0])*100+int(word[1])] = 1
	}
	return ans
}
