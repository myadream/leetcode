package decode

//https://leetcode-cn.com/problems/decode-xored-array/
func decode(encoded []int, first int) []int {
	var n = len(encoded)
	var ans = make([]int, n + 1)

	ans[0] = first

	for i := 0; i < n; i++ {
		ans[i + 1] = ans[i] ^ encoded[i]
	}

	return ans
}