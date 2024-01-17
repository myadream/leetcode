package leetcode_11

import "leetcode/src/common"

func caseOne(height []int) (ans int) {
	n := len(height)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			sum := (j - i) * common.Min(height[i], height[j])
			ans = common.Max(ans, sum)

			if height[i] < height[j] {
				break
			}
		}
	}

	return ans
}

func caseTwo(height []int) (ans int) {
	l, r := 0, len(height)-1

	for l < r {
		sum := (r - l) * common.Min(height[l], height[r])
		ans = common.Max(ans, sum)

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return ans
}
