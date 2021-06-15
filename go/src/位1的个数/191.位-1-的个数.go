

/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) (ans int) {

	for i := 0; i <= 31 && num > 0; i++ {
		if num&1 == 1 {
			ans++
		}

		num >>= 1
	}

	return
}

// @lc code=end
