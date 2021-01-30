package findNumbers

//https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
func findNumbers(nums []int) int {
	var res int

	for i := 0; i < len(nums); i++ {
		t := 0

		for nums[i] != 0 {
			nums[i] /= 10
			t++
		}

		if t % 2 == 0 {
			res++
		}
	}

	return res
}