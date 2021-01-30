package kidsWithCandies

//https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var n, max  = len(candies), 0
	var res = make([]bool, n)

	for i := 0; i < n; i++ {
		if max < candies[i] {
			max = candies[i]
		}
	}

	for i := 0; i < n; i++ {
		if candies[i] + extraCandies >= max {
			res[i] = true
		}
	}

	return res
}