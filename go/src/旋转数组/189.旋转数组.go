/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])

}

func reverse(arr []int) {
	var l = len(arr)
	for i := 0; i < l / 2; i++ {
		arr[i], arr[l - i - 1] = arr[l - i - 1], arr[i]
	}
}
// @lc code=end

