package hammingWeight

//https://leetcode-cn.com/problems/er-jin-zhi-zhong-1de-ge-shu-lcof/solution/mian-shi-ti-15-er-jin-zhi-zhong-1de-ge-shu-wei-yun/
func hammingWeight(num uint32) int {
	var ans uint32

	for num > 0 {
		ans += num & 1
		num >>= 1
	}

	return int(ans)
}