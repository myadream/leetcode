package arrayPairSum

import "sort"

//https://leetcode-cn.com/problems/array-partition-i/
func arrayPairSum(nums []int) (ans int) {
	sort.Ints(nums)
	var n= len(nums)

	for i := 0; i < n; i += 2 {
		ans += nums[i]
	}

	return
}