package template

func caseOne(n int) int {
	var fn func(i int) int
	var dp = make(map[int]int, n+1)

	fn = func(i int) int {
		if i <= 1 {
			return 1
		}

		if dp[i] > 0 {
			return dp[i]
		}

		dp[i] = fn(i-1) + fn(i-2)

		return dp[i]
	}

	return fn(n-1) + fn(n-2)
}
