package leetcode_2765

import "leetcode/src/common"

func caseOne(nums []int) int {

	n := len(nums)
	ans := 0

	for i := 0; i < n-1; {
		if nums[i+1]-nums[i] != 1 {
			i++
			continue
		}

		i0 := i
		i += 2

		for i < n && nums[i] == nums[i-2] {
			i++
		}

		ans = common.Max(ans, i-i0)
		i -= 1
	}

	return ans
}

func caseTwo(nums []int) int {
	n := len(nums)
	ans := -1

	for i := 0; i < n; i++ {

		flag := 1
		for j := i + 1; j < n; {
			if nums[j]-nums[j-1] != flag {
				break
			}

			j++
			flag = -(flag)
			ans = common.Max(ans, j-i)
		}
	}

	return ans
}
