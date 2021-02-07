package checkPossibility

//https://leetcode-cn.com/problems/non-decreasing-array/
func checkPossibility(nums []int) bool {
	var count = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i - 1] {
			if i == 1 || nums[i] >= nums[i - 2] {
				nums[i - 1] = nums[i]
			} else {
				nums[i] = nums[i - 1]
			}

			count++
		}
	}

	return count <= 1
}
