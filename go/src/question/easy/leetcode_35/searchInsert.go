package leetcode_35

// 默认搜索
func defaultSearch(nums []int, target int) int {
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			return i
		}
	}

	return len(nums)
}

func binarySearch(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		mid := l + (r - l) - 1
		if nums[mid] >= target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}
