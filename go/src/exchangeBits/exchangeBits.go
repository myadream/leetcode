package exchangeBits

//https://leetcode-cn.com/problems/exchange-lcci/
func exchangeBits(num int) int {
	var even = num &  0x55555555
	var odd = num & 0xaaaaaaaa

	even = even << 1
	odd = odd >> 1

	return even | odd
}
