package minTimeToVisitAllPoints

//https://leetcode-cn.com/problems/minimum-time-visiting-all-points/
func minTimeToVisitAllPoints(points [][]int) int {
	var ans = 0
	var abs = func(n int) int {
		if n < 0 {
			return ^n + 1
		}

		return n
	}

	var max = func(x, y int) int {
		if x > y {
			return x
		}

		return y
	}

	for i := 1 ;i < len(points); i++ {
		ans += max(abs(points[0][0] - points[i][0]), abs(points[0][1] - points[i][1]))
		points[0][0] = points[i][0]
		points[0][1] = points[i][1]
	}

	return ans
}