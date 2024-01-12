package leetcode_2645

func caseOne(word string) int {
	l := len(word)
	dp := make([]int, l+1)

	for i := 1; i <= l; i++ {
		dp[i] = dp[i-1] + 2
		if i > 1 && word[i-1] > word[i-2] {
			dp[i] = dp[i-1] - 1
		}
	}

	return dp[l]
}

func caseTwo(word string) int {
	var l = len(word)

	ans := 1
	for i := 1; i < l; i++ {
		if word[i] <= word[i-1] {
			ans += 1
		}
	}

	return ans*3 - l
}
