package leetcode_746

import "leetcode/src/common"

func caseOne(cost []int) int {
	var dfs func(int) int
	dfs = func(i int) int {
		if i <= 1 { // 递归边界
			return 0
		}
		return common.Min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])
	}
	return dfs(len(cost))
}

func caseTwo(cost []int) int {
	l := len(cost)
	var dfs func(int) int
	var dp = make([]int, l+1)

	dfs = func(i int) int {
		if i <= 1 {
			return 0
		}

		if dp[i] != 0 {
			return dp[i]
		}

		dp[i] = common.Min(dfs(i-1)+cost[i-1], dfs(i-2)+cost[i-2])

		return dp[i]
	}

	return dfs(l)
}

func caseThree(cost []int) int {
	f1, f2 := 0, 0
	for i := 1; i < len(cost); i++ {
		f1, f2 = f2, common.Min(f2+cost[i], f1+cost[i-1])
	}

	return f2
}
