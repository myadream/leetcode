package longestOnes

//https://leetcode-cn.com/problems/max-consecutive-ones-iii/
func longestOnes(A []int, K int) (ans int) {
	var left, right, zeros int
	var n = len(A) - 1

	for right <= n {
		if A[right] == 0 {
			zeros++
		}

		for zeros > K {
			if A[left] == 0 {
				zeros -= 1
			}

			left += 1
		}

		ans = max(ans, right - left + 1)
		right++
	}

	return
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}