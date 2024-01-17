package leetcode_283

func caseOne(nums []int) []int {

	l, r := 0, 1
	n := len(nums) - 1

	for r <= n {
		if nums[r] != 0 {
			nums[r], nums[l] = nums[l], nums[r]
			l++
		}

		r++
	}

	return nums
}

func caseTwo(nums []int) []int {
	var j = 0
	var n = len(nums)

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for j < n {
		nums[j] = 0
		j++
	}

	return nums
}
