package template

import "math/bits"

func countBits(n int) int {
	count := 0
	for n != 0 {
		n &= n - 1
		count++
	}
	return count
}

func caseOne(nums []int, k int) int {

	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if countBits(i) == k {
			ans += nums[i]
		}
	}

	return ans
}

func caseTwo(nums []int, k int) int {

	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if bits.OnesCount(uint(i)) == k {
			ans += nums[i]
		}
	}

	return ans
}
