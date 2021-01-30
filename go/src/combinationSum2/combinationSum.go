package combinationSum

import (
	"sort"
)

var res [][]int

func combinationSum(candidates []int, target int) [][]int {
	res = res[:0]
	sort.Ints(candidates)

	dfs([]int(nil), candidates, 0, len(candidates), target)

	return res
}

func dfs(path, candidates []int, begin, lens, target int) {
	if target == 0 {
		res = append(res, append([]int(nil), path...))
	}

	for i := begin; i < lens; i++ {
		if target - candidates[i] < 0 {
			break
		}

		if i > begin && candidates[i] == candidates[i - 1] {
			continue
		}

		dfs(append(path, candidates[i]), candidates, i + 1, lens, target - candidates[i])
	}
}