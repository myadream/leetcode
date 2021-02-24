package singleNumber

//https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x21ib6/
func singleNumber(nums []int) (ans int) {
	for i := 0; i < len(nums); i++ {
		ans = ans ^ nums[i]
	}

	return
}
