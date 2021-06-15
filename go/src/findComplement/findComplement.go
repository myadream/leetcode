package findComplement

//https://leetcode-cn.com/problems/number-complement/comments/
func findComplement(num int) int {
	var ans int
	for ans < num {
		ans = (ans << 1) + 1
	}

	return ans ^ num
}