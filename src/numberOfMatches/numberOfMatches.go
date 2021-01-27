package numberOfMatches

//https://leetcode-cn.com/problems/count-of-matches-in-tournament/
func numberOfMatches(n int) int {
	var res int

	for n > 1 {
		res += n >> 1
		n = (n >> 1) + (n & 1)

	}

	return res
}