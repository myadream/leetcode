package firstMissingPositive

//https://leetcode-cn.com/problems/first-missing-positive/solution/que-shi-de-di-yi-ge-zheng-shu-by-leetcode-solution/

//哈希数据表
func firstMissingPositiveOne(nums []int) int {
	var n = len(nums)

	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	for i := 0; i < n; i++ {
		num := abs(nums[i])

		if num <= n {
			nums[num - 1] = ^abs(nums[num - 1]) + 1
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}

func abs(num int) int {
	if num < 0 {
		return ^num + 1
	}

	return num
}