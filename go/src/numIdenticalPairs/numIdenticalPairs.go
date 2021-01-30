package numIdenticalPairs

//https://leetcode-cn.com/problems/number-of-good-pairs/
func numIdenticalPairs(nums []int) int {
	var sum, n = 0, len(nums)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i] == nums[j] {
				sum++
			}
		}
	}

	return sum
}
