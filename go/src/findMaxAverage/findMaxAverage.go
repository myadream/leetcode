package findMaxAverage

//https://leetcode-cn.com/problems/maximum-average-subarray-i/solution/zi-shu-zu-zui-da-ping-jun-shu-i-by-leetc-us1k/
func findMaxAverage(nums []int, k int) float64 {
	var sum, n = 0, len(nums)

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	maxSum := sum
	for i := k; i < n; i++ {
		sum = sum - nums[i - k] + nums[i]
		maxSum = max(maxSum, sum)
	}

	return float64(maxSum) / float64(k)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
