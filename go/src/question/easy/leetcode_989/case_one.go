package leetcode_989

import "leetcode/src/common"

func CaseOne(num []int, k int) []int {
	sum := 0
	var total = make([]int, 0)
	for l := len(num); l > 0 || sum == 1 || k > 0; {
		if l > 0 {
			l--
			sum += num[l]
		}

		sum += k % 10

		total = append(total, sum%10)
		sum /= 10
		k /= 10
	}

	return common.Reverse(total)
}
