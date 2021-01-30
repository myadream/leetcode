package maxArea

func maxArea(height []int) int {
	var area, l, r = 0, 0, len(height) - 1

	for l < r {
		ans := (r - l) * min(height[l], height[r])

		if ans > area {
			area = ans
		}

		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}

	return area
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}