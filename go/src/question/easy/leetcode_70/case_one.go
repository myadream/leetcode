package leetcode_70

func caseOne(n int) int {

	var dfs func(n int) int
	dfs = func(n int) int {
		if n <= 1 {
			return 1
		}

		return dfs(n-1) + dfs(n-2)
	}

	return dfs(n-1) + dfs(n-2)
}

func caseTwo(n int) int {
	var dfs func(n int) int
	var dp = make([]int, n+1)
	dfs = func(n int) int {
		if n <= 1 {
			return 1
		}

		if dp[n] != 0 {
			return dp[n]
		}

		dp[n] = dfs(n-1) + dfs(n-2)

		return dp[n]
	}

	return dfs(n-1) + dfs(n-2)
}

func caseThree(n int) int {
	var dp = make([]int, n+1)

	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}

func caseFour(n int) int {
	f1, f2, ans := 0, 0, 1
	for i := 0; i < n; i++ {
		f1 = f2
		f2 = ans
		ans = f1 + f2
	}

	return ans
}
