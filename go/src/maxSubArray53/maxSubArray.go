package maxSubArray53

func maxSubArray(nums []int) int {
	var (
		maxCount = nums[0]
		count    = nums[0]
	)

	for i := 1; i < len(nums); i++ {
		count = max(count+nums[i], nums[i])
		maxCount = max(maxCount, count)
	}

	return maxCount
}

func max(l, r int) int {
	if l > r {
		return l
	} else {
		return r
	}
}
