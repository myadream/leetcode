package xorOperation

//https://leetcode-cn.com/problems/xor-operation-in-an-array/
func xorOperation(n int, start int) int {
	var res = 0

	for n != 0 {
		n--
		res ^= start + 2 * n
	}

	return res
}