package numberOfSteps

//https://leetcode-cn.com/problems/number-of-steps-to-reduce-a-number-to-zero/
func numberOfSteps (num int) (ans int) {
	for num | 0 != 0 {
		if num & 1 == 1 {
			num &= -2
		} else {
			num >>= 1
		}

		ans++
	}

	return ans
}