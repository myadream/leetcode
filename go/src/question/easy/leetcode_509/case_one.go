package leetcode_509

func caseOne(n int) int {
	var dfs func(i int) int
	var dp = make(map[int]int, n+1)

	dfs = func(i int) int {
		if i <= 0 {
			return 0
		}

		if i <= 1 {
			return 1
		}

		if dp[i] != 0 {
			return dp[i]
		}

		dp[i] = dfs(i-1) + dfs(i-2)

		return dp[i]
	}

	return dfs(n)
}

func caseTwo(n int) int {

	f1, f2 := 0, 1
	for i := 0; i < n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f1
}
func caseThree(n int) int {
	if n <= 1 {
		return n
	}

	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}

	return f2
}
