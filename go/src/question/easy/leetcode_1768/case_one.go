package leetcode_1768

func caseOne(word1 string, word2 string) string {
	w1, w2, i, j := len(word1)-1, len(word2)-1, 0, 0

	ans := ""
	for i <= w1 || j <= w2 {
		if i <= w1 {
			ans += string(word1[i])
			i += 1
		}

		if j <= w2 {
			ans += string(word2[j])
			j += 1
		}
	}

	return ans
}

func caseTwo(word1 string, word2 string) string {
	w1, w2 := len(word1)-1, len(word2)-1

	var ans = make([]byte, 0, w1+w2)
	for i := 0; i <= w1 || i <= w2; i++ {
		if i <= w1 {
			ans = append(ans, word1[i])
		}

		if i <= w2 {
			ans = append(ans, word2[i])
		}
	}

	return string(ans)
}
