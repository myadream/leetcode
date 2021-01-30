package threeSum

import (
	"sort"
)
/**
func threeSum(nums []int) [][]int {

	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)
	numsLength := len(nums)

	var res [][]int
	for k := 0; k < numsLength - 2; k ++ {

		if nums[k] > 0 {
			break
		}

		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		left  := k + 1
		right := numsLength - 1

		for left < right {

			sum := nums[k] + nums[left] + nums[right]
			if sum > 0 {
				for right = right - 1; left < right && nums[right] == nums[right+1]; right -- {}
			}else if sum < 0 {
				for left = left + 1; left < right && nums[left] == nums[left-1]; left ++ {}
			}else {
				res = append(res, []int{nums[k], nums[left], nums[right]})
				for left = left + 1; left < right && nums[left] == nums[left-1]; left ++ {}
				for right = right - 1; left < right && nums[right] == nums[right+1]; right -- {}
			}
		}
	}

	return res
}
 */

//三数之和
func threeSum(nums []int) [][]int {
	var n, i, R, L int
	var res [][]int

	n = len(nums)
	if n < 3 {
		return nil
	}

	sort.Ints(nums)

	for i = 0; i < n; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		L = i + 1
		R = n - 1

		for L < R {
			t := nums[i] + nums[R] + nums[L]
			if t == 0 {
				res = append(res, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L + 1] {
					L++
				}

				for L < R && nums[R] == nums[R - 1] {
					R--
				}

				L++
				R--
			} else if t > 0 {
				R--
			} else {
				L++
			}
		}
	}

	return res
}