package leetcode_15

import "sort"

func caseOne(nums []int) [][]int {

	sort.Ints(nums)

	n := len(nums)
	var ans = make([][]int, 0)
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ans
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, n-1

		for l < r {
			if nums[l]+nums[r] > -nums[i] {
				r--
			} else if nums[l]+nums[r] < -nums[i] {
				l++
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				for l++; l < r && nums[l] == nums[l-1]; l++ {
				}

				for r--; r > l && nums[r] == nums[r+1]; r-- {
				}
			}
		}
	}

	return ans
}
