package threeSumClosest

import (
	"sort"
)

//最接近的三数之和
//https://leetcode-cn.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	var n, best, i, L, R, sum int

	best = 1e7
	n = len(nums) - 1
	sort.Ints(nums)

	for i = 0; i < n; i++ {
		L = i + 1
		R = n

		for L < R {
			sum = nums[i] + nums[L] + nums[R]

			if sum == target {
				return target
			}

			if abs(sum - target) < abs(best - target) {
				best = sum
			}

			if sum >= target {
				for L < R &&  nums[R] == nums[R - 1] {
					R--
				}

				R--
			} else {
				for L < R && nums[L] == nums[L + 1] {
					L++
				}

				L++
			}
		}
	}

	return best
}

func abs(num int) int {
	if num > 0 {
		return num
	}

	return ^num + 1
}
