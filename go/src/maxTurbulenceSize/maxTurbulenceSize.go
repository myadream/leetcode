package maxTurbulenceSize

//https://leetcode-cn.com/problems/longest-turbulent-subarray/
func maxTurbulenceSize(arr []int) int {
	var left, ans, n = 0, 1, len(arr)

	for right := 1; right < n; right++ {
		if arr[right] == arr[right - 1] {
			left = right
		} else if right != 1 && (arr[right] - arr[right - 1]) ^ (arr[right - 1] - arr[right - 2]) >= 0 {
			left = right - 1
		}

		ans = max(ans, right - left + 1)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}