package search

//https://leetcode-cn.com/problems/search-in-rotated-sorted-array/comments/
func search(nums []int, target int) int {
	var l, r, mid = 0, len(nums) - 1, 0

	for l <= r {
		mid = l + ((r - l) / 2)

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < nums[r] {
			if nums[mid] < target && nums[r] >= target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		} else {
			if nums[l] <= target && nums[mid] > target {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}
	}

	return -1
}