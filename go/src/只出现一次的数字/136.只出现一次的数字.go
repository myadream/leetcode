/*
 * @lc app=leetcode.cn id=136 lang=golang
 *
 * [136] 只出现一次的数字
 */

// @lc code=start
func singleNumber(nums []int) (ans int) {
	for _, v := range nums {
		ans ^= v
	}

	return
}

// @lc code=end
