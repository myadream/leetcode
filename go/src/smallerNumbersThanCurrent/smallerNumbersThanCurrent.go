package smallerNumbersThanCurrent

//https://leetcode-cn.com/problems/how-many-numbers-are-smaller-than-the-current-number/
//暴力
func smallerNumbersThanCurrentOne(nums []int) []int {
	var ans = make([]int, len(nums))

	for k, val := range nums {
		for _, v := range nums {
			if val > v {
				ans[k]++
			}
		}
	}

	return ans
}

//map
func smallerNumbersThanCurrentTwo(nums []int) []int {
	var n = len(nums)
	var ans = make([]int, 101)

	for i := 0; i < n; i++ {
		ans[nums[i]]++
	}

	for i := 1; i < 101; i++ {
		ans[i] += ans[i - 1]
	}


	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			nums[i] = 0
		} else {
			nums[i] = ans[nums[i] - 1]
		}
	}

	return nums
}