package maxProfit

//https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii/solution/mai-mai-gu-piao-de-zui-jia-shi-ji-ii-by-leetcode-s/
func maxProfit(prices []int) int {
	var ans, n = 0, len(prices)

	for i := 1; i < n; i++ {
		ans += max(0, prices[i] - prices[i - 1])
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}