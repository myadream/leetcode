package combinationSum

import "sort"

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = res[:0]
	sort.Ints(candidates)

	dfs([]int(nil), candidates, target)

	return res
}

func dfs(path, candidates []int, target int) {
	if target == 0 {
		res = append(res, append([]int(nil), path...))
	}

	for i, v := range candidates {
		if target - v < 0 {
			return
		}

		dfs(append(path, v), candidates[i:], target - v)
	}
}