
/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	var m = map[int]int{}
	var k = 0

	//map
	for _, v := range nums1 {
		m[v]++
	}

	for _, v := range nums2 {
		if m[v] <= 0 {
			continue
		}

		nums1[k] = v
		k++
		m[v]--
	}

	return nums1[:k]
}

// @lc code=end
