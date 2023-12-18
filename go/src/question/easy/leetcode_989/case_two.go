package leetcode_989

import "leetcode/src/common"

func CaseTwo(num []int, k int) []int {
	var total = make([]int, 0)
	for l := len(num) - 1; l >= 0 || k > 0; l-- {
		if l >= 0 {
			k += num[l]
		}

		total = append(total, k%10)
		k /= 10
	}

	return common.Reverse(total)
}
