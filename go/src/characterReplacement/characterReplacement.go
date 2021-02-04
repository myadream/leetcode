package characterReplacement

//https://leetcode-cn.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {
	var maxCht, left = 0, 0
	var cnt = make([]int, 26)

	for right, ch := range s {
		cnt[ch - 'A']++

		maxCht = max(maxCht, cnt[ch-'A'])

		if right - left + 1 - maxCht > k {
			cnt[s[left]-'A']--
			left++
		}
	}

	return len(s) - left
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}