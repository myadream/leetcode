package removeDuplicates

//https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/solution/
func removeDuplicates(nums []int) int {
	var res int

	for i, _ := range nums {
		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		if nums[res] != nums[i] {
			res++
			nums[res] = nums[i]
		}
	}

	return res + 1
}
