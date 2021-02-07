package maxScore

//https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/
func maxScore(cardPoints []int, k int) int {
	var n, sum = len(cardPoints), 0
	var windowSize = n - k

	for _, v := range cardPoints[:windowSize] {
		sum += v
	}

	total := sum
	minSum := sum
	for i := windowSize; i < n; i++ {
		total += cardPoints[i]
		sum += cardPoints[i] - cardPoints[i - windowSize]
		minSum = min(minSum, sum)
	}

	return total - minSum
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
