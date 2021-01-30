package searchInsert

//https://leetcode-cn.com/problems/search-insert-position/
func searchInsert(nums []int, target int) int {
	var L, R, mid = 0, len(nums) - 1, 0

	for L <= R {
		mid = L + ((R - L) >> 1)

		if nums[mid] < target {
			L = mid + 1
		} else {
			R = mid - 1
		}
	}

	return L
}
