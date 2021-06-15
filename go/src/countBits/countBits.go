package countBits

//https://leetcode-cn.com/problems/counting-bits/
func countBits(num int) (ans []int) {
	ans = make([]int, num + 1)

	for i := 1; i <= num; i++ {
		ans[i] = ans[i & (i - 1)] + 1
	}

	return
}

//plan one
/**
func countBits(num int) (ans []int) {
	ans = make([]int, num + 1)

	for i := 1; i <= num; i++ {
		if i & 1 == 1 {
			ans[i] = ans[i - 1] + 1
		} else 	{
			ans[i] = ans[i / 2]
		}
	}

	return
}
 */